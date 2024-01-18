package querybuilder

import (
	"errors"
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"strings"
)

type bindingType string
type Builder[T any] struct {
	contracts.QueryExecutor[T]
	limit    int64
	offset   int64
	distinct bool
	table    string
	fields   []string
	wheres   *Wheres
	orderBy  OrderByFields
	groupBy  GroupBy
	joins    Joins
	unions   Unions[T]
	having   *Wheres
	bindings map[bindingType][]any
}

func (builder *Builder[T]) Bind(executor contracts.QueryExecutor[T]) contracts.Query[T] {
	builder.QueryExecutor = executor
	return builder
}

func (builder *Builder[T]) Skip(offset int64) contracts.Query[T] {
	return builder.Offset(offset)
}

func (builder *Builder[T]) Take(num int64) contracts.Query[T] {
	return builder.Limit(num)
}

const (
	selectBinding  bindingType = "select"
	fromBinding    bindingType = "from"
	joinBinding    bindingType = "join"
	whereBinding   bindingType = "where"
	groupByBinding bindingType = "groupBy"
	havingBinding  bindingType = "having"
	orderBinding   bindingType = "order"
	unionBinding   bindingType = "union"
)

func (builder *Builder[T]) getWheres() *Wheres {
	return builder.wheres
}

func (builder *Builder[T]) prepareArgs(condition string, args any) (raw string, bindings []any) {
	if expression, isExpression := args.(Expression); isExpression {
		return string(expression), bindings
	} else if instance, isBuilder := args.(contracts.QueryBuilder[T]); isBuilder {
		raw, bindings = instance.SelectSql()
		raw = fmt.Sprintf("(%s)", raw)
		return
	}
	condition = strings.ToLower(condition)
	switch condition {
	case "in", "not in", "between", "not between":
		isInGrammar := strings.Contains(condition, "in")
		joinSymbol := utils.IfString(isInGrammar, ",", " and ")
		var stringArg string
		switch arg := args.(type) {
		case string:
			stringArg = arg
		case fmt.Stringer:
			stringArg = arg.String()
		case []string:
			stringArg = strings.Join(arg, joinSymbol)
		case []int:
			stringArg = utils.JoinIntArray(arg, joinSymbol)
		case []int64:
			stringArg = utils.JoinInt64Array(arg, joinSymbol)
		case []float64:
			stringArg = utils.JoinFloat64Array(arg, joinSymbol)
		case []float32:
			stringArg = utils.JoinFloatArray(arg, joinSymbol)
		case []any:
			bindings = arg
			raw = fmt.Sprintf("(%s)", strings.Join(utils.MakeSymbolArray("?", len(bindings)), ","))
			return
		default:
			panic(ParamException{
				Err:       errors.New("不支持的参数类型"),
				Arg:       arg,
				Condition: condition,
			})
		}
		bindings = utils.StringArray2InterfaceArray(strings.Split(stringArg, joinSymbol))
		if isInGrammar {
			raw = fmt.Sprintf("(%s)", strings.Join(utils.MakeSymbolArray("?", len(bindings)), ","))
		} else {
			raw = "? and ?"
		}
	case "is", "is not", "exists", "not exists":
		raw = utils.ToString(args, "")
	default:
		raw = "?"
		bindings = append(bindings, utils.ToString(args, ""))
	}

	return
}

func (builder *Builder[T]) addBinding(bindType bindingType, bindings ...any) contracts.Query[T] {
	builder.bindings[bindType] = append(builder.bindings[bindType], bindings...)
	return builder
}

func (builder *Builder[T]) GetBindings() (results []any) {
	for _, binding := range []bindingType{
		selectBinding, fromBinding, joinBinding,
		whereBinding, groupByBinding, havingBinding, orderBinding, unionBinding,
	} {
		results = append(results, builder.bindings[binding]...)
	}
	return
}

func (builder *Builder[T]) Distinct() contracts.Query[T] {
	builder.distinct = true
	return builder
}

func (builder *Builder[T]) From(table string, as ...string) contracts.Query[T] {
	if len(as) == 0 {
		builder.table = table
	} else {
		builder.table = fmt.Sprintf("%s as %s", table, as[0])
	}
	return builder
}

func (builder *Builder[T]) Offset(offset int64) contracts.Query[T] {
	builder.offset = offset
	return builder
}

func (builder *Builder[T]) Limit(num int64) contracts.Query[T] {
	builder.limit = num
	return builder
}

func (builder *Builder[T]) WithPagination(perPage int64, current ...int64) contracts.Query[T] {
	builder.limit = perPage
	if len(current) > 0 {
		builder.offset = perPage * (current[0] - 1)
	}
	return builder
}

func (builder *Builder[T]) FromMany(tables ...string) contracts.Query[T] {
	if len(tables) > 0 {
		builder.table = strings.Join(tables, ",")
	}
	return builder
}

func (builder *Builder[T]) FromSub(provider contracts.QueryProvider[T], as string) contracts.Query[T] {
	subBuilder := provider()
	builder.table = fmt.Sprintf("(%s) as %s", subBuilder.ToSql(), as)
	return builder.addBinding(fromBinding, subBuilder.GetBindings()...)
}

func (builder *Builder[T]) When(condition bool, callback contracts.QueryCallback[T], elseCallback ...contracts.QueryCallback[T]) contracts.Query[T] {
	if condition {
		return callback(builder)
	} else if len(elseCallback) > 0 {
		return elseCallback[0](builder)
	}
	return builder
}

func (builder *Builder[T]) getSelect() string {
	if builder.distinct {
		return "distinct " + strings.Join(builder.fields, ",")
	}
	return strings.Join(builder.fields, ",")
}

func (builder *Builder[T]) ToSql() string {
	sql := fmt.Sprintf("select %s from `%s`", builder.getSelect(), builder.table)

	if !builder.joins.IsEmpty() {
		sql = fmt.Sprintf("%s %s", sql, builder.joins.String())
	}

	if !builder.wheres.IsEmpty() {
		sql = fmt.Sprintf("%s where %s", sql, builder.wheres.String())
	}

	if !builder.groupBy.IsEmpty() {
		sql = fmt.Sprintf("%s group by %s", sql, builder.groupBy.String())

		if !builder.having.IsEmpty() {
			sql = fmt.Sprintf("%s having %s", sql, builder.having.String())
		}
	}

	if !builder.orderBy.IsEmpty() {
		sql = fmt.Sprintf("%s order by %s", sql, builder.orderBy.String())
	}

	if builder.limit > 0 {
		sql = fmt.Sprintf("%s limit %d", sql, builder.limit)
	}
	if builder.offset > 0 {
		sql = fmt.Sprintf("%s offset %d", sql, builder.offset)
	}

	if !builder.unions.IsEmpty() {
		sql = fmt.Sprintf("(%s) %s", sql, builder.unions.String())
	}

	return sql
}

func (builder *Builder[T]) SelectSql() (string, []any) {
	return builder.ToSql(), builder.GetBindings()
}

func (builder *Builder[T]) SelectForUpdateSql() (string, []any) {
	return builder.ToSql() + " for update", builder.GetBindings()
}
