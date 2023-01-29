package querybuilder

import (
	"errors"
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"strings"
)

type bindingType string
type Builder struct {
	contracts.QueryBuilder
	limit    int64
	offset   int64
	distinct bool
	table    string
	fields   []string
	wheres   *Wheres
	orderBy  OrderByFields
	groupBy  GroupBy
	joins    Joins
	unions   Unions
	having   *Wheres
	bindings map[bindingType][]interface{}
}

func (builder *Builder) Bind(b contracts.QueryBuilder) contracts.QueryBuilder {
	builder.QueryBuilder = b
	return builder
}

func (builder *Builder) Skip(offset int64) contracts.QueryBuilder {
	return builder.Offset(offset)
}

func (builder *Builder) Take(num int64) contracts.QueryBuilder {
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

func NewQuery(table string) *Builder {
	return &Builder{
		table:    table,
		fields:   []string{"*"},
		orderBy:  OrderByFields{},
		bindings: map[bindingType][]interface{}{},
		joins:    Joins{},
		unions:   Unions{},
		groupBy:  GroupBy{},
		wheres: &Wheres{
			wheres:    map[contracts.WhereJoinType][]*Where{},
			subWheres: map[contracts.WhereJoinType][]*Wheres{},
		},
		having: &Wheres{
			wheres:    map[contracts.WhereJoinType][]*Where{},
			subWheres: map[contracts.WhereJoinType][]*Wheres{},
		},
	}
}

func FromSub(callback contracts.QueryProvider, as string) contracts.QueryBuilder {
	return NewQuery("").FromSub(callback, as)
}

func (builder *Builder) getWheres() *Wheres {
	return builder.wheres
}

func (builder *Builder) prepareArgs(condition string, args interface{}) (raw string, bindings []interface{}) {
	if expression, isExpression := args.(Expression); isExpression {
		return string(expression), bindings
	} else if builder, isBuilder := args.(contracts.QueryBuilder); isBuilder {
		raw, bindings = builder.SelectSql()
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
		case []interface{}:
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
		raw = utils.ConvertToString(args, "")
	default:
		raw = "?"
		bindings = append(bindings, utils.ConvertToString(args, ""))
	}

	return
}

func (builder *Builder) addBinding(bindType bindingType, bindings ...interface{}) contracts.QueryBuilder {
	builder.bindings[bindType] = append(builder.bindings[bindType], bindings...)
	return builder
}

func (builder *Builder) GetBindings() (results []interface{}) {
	for _, binding := range []bindingType{
		selectBinding, fromBinding, joinBinding,
		whereBinding, groupByBinding, havingBinding, orderBinding, unionBinding,
	} {
		results = append(results, builder.bindings[binding]...)
	}
	return
}

func (builder *Builder) Distinct() contracts.QueryBuilder {
	builder.distinct = true
	return builder
}

func (builder *Builder) From(table string, as ...string) contracts.QueryBuilder {
	if len(as) == 0 {
		builder.table = table
	} else {
		builder.table = fmt.Sprintf("%s as %s", table, as[0])
	}
	return builder
}

func (builder *Builder) Offset(offset int64) contracts.QueryBuilder {
	builder.offset = offset
	return builder
}

func (builder *Builder) Limit(num int64) contracts.QueryBuilder {
	builder.limit = num
	return builder
}

func (builder *Builder) WithPagination(perPage int64, current ...int64) contracts.QueryBuilder {
	builder.limit = perPage
	if len(current) > 0 {
		builder.offset = perPage * (current[0] - 1)
	}
	return builder
}

func (builder *Builder) FromMany(tables ...string) contracts.QueryBuilder {
	if len(tables) > 0 {
		builder.table = strings.Join(tables, ",")
	}
	return builder
}

func (builder *Builder) FromSub(provider contracts.QueryProvider, as string) contracts.QueryBuilder {
	subBuilder := provider()
	builder.table = fmt.Sprintf("(%s) as %s", subBuilder.ToSql(), as)
	return builder.addBinding(fromBinding, subBuilder.GetBindings()...)
}

func (builder *Builder) When(condition bool, callback contracts.QueryCallback, elseCallback ...contracts.QueryCallback) contracts.QueryBuilder {
	if condition {
		return callback(builder)
	} else if len(elseCallback) > 0 {
		return elseCallback[0](builder)
	}
	return builder
}

func (builder *Builder) getSelect() string {
	if builder.distinct {
		return "distinct " + strings.Join(builder.fields, ",")
	}
	return strings.Join(builder.fields, ",")
}

func (builder *Builder) ToSql() string {
	sql := fmt.Sprintf("select %s from %s", builder.getSelect(), builder.table)

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

func (builder *Builder) SelectSql() (string, []interface{}) {
	return builder.ToSql(), builder.GetBindings()
}

func (builder *Builder) SelectForUpdateSql() (string, []interface{}) {
	return builder.ToSql() + " for update", builder.GetBindings()
}
