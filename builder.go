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

func (this *Builder) Bind(builder contracts.QueryBuilder) contracts.QueryBuilder {
	this.QueryBuilder = builder
	return this
}

func (this *Builder) Skip(offset int64) contracts.QueryBuilder {
	return this.Offset(offset)
}

func (this *Builder) Take(num int64) contracts.QueryBuilder {
	return this.Limit(num)
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

func (this *Builder) getWheres() *Wheres {
	return this.wheres
}

func (this *Builder) prepareArgs(condition string, args interface{}) (raw string, bindings []interface{}) {
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
			panic(ParamException{errors.New("不支持的参数类型"), contracts.Fields{
				"arg":       arg,
				"condition": condition,
			}})
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

func (this *Builder) addBinding(bindType bindingType, bindings ...interface{}) contracts.QueryBuilder {
	this.bindings[bindType] = append(this.bindings[bindType], bindings...)
	return this
}

func (this *Builder) GetBindings() (results []interface{}) {
	for _, binding := range []bindingType{
		selectBinding, fromBinding, joinBinding,
		whereBinding, groupByBinding, havingBinding, orderBinding, unionBinding,
	} {
		results = append(results, this.bindings[binding]...)
	}
	return
}

func (this *Builder) Distinct() contracts.QueryBuilder {
	this.distinct = true
	return this
}

func (this *Builder) From(table string, as ...string) contracts.QueryBuilder {
	if len(as) == 0 {
		this.table = table
	} else {
		this.table = fmt.Sprintf("%s as %s", table, as[0])
	}
	return this
}

func (this *Builder) Offset(offset int64) contracts.QueryBuilder {
	this.offset = offset
	return this
}

func (this *Builder) Limit(num int64) contracts.QueryBuilder {
	this.limit = num
	return this
}

func (this *Builder) WithPagination(perPage int64, current ...int64) contracts.QueryBuilder {
	this.limit = perPage
	if len(current) > 0 {
		this.offset = perPage * (current[0] - 1)
	}
	return this
}

func (this *Builder) FromMany(tables ...string) contracts.QueryBuilder {
	if len(tables) > 0 {
		this.table = strings.Join(tables, ",")
	}
	return this
}

func (this *Builder) FromSub(provider contracts.QueryProvider, as string) contracts.QueryBuilder {
	subBuilder := provider()
	this.table = fmt.Sprintf("(%s) as %s", subBuilder.ToSql(), as)
	return this.addBinding(fromBinding, subBuilder.GetBindings()...)
}

func (this *Builder) When(condition bool, callback contracts.QueryCallback, elseCallback ...contracts.QueryCallback) contracts.QueryBuilder {
	if condition {
		return callback(this)
	} else if len(elseCallback) > 0 {
		return elseCallback[0](this)
	}
	return this
}

func (this *Builder) getSelect() string {
	if this.distinct {
		return "distinct " + strings.Join(this.fields, ",")
	}
	return strings.Join(this.fields, ",")
}

func (this *Builder) ToSql() string {
	sql := fmt.Sprintf("select %s from %s", this.getSelect(), this.table)

	if !this.joins.IsEmpty() {
		sql = fmt.Sprintf("%s %s", sql, this.joins.String())
	}

	if !this.wheres.IsEmpty() {
		sql = fmt.Sprintf("%s where %s", sql, this.wheres.String())
	}

	if !this.groupBy.IsEmpty() {
		sql = fmt.Sprintf("%s group by %s", sql, this.groupBy.String())

		if !this.having.IsEmpty() {
			sql = fmt.Sprintf("%s having %s", sql, this.having.String())
		}
	}

	if !this.orderBy.IsEmpty() {
		sql = fmt.Sprintf("%s order by %s", sql, this.orderBy.String())
	}

	if this.limit > 0 {
		sql = fmt.Sprintf("%s limit %d", sql, this.limit)
	}
	if this.offset > 0 {
		sql = fmt.Sprintf("%s offset %d", sql, this.offset)
	}

	if !this.unions.IsEmpty() {
		sql = fmt.Sprintf("(%s) %s", sql, this.unions.String())
	}

	return sql
}

func (this *Builder) SelectSql() (string, []interface{}) {
	return this.ToSql(), this.GetBindings()
}

func (this *Builder) SelectForUpdateSql() (string, []interface{}) {
	return this.ToSql() + " for update", this.GetBindings()
}
