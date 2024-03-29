package querybuilder

import (
	"fmt"
	"github.com/goal-web/contracts"
)

type Where struct {
	field     string
	condition string
	arg       string
}

func (this *Where) String() string {
	if this == nil {
		return ""
	}
	return fmt.Sprintf("%s %s %s", this.field, this.condition, this.arg)
}

type Wheres struct {
	subWheres map[contracts.WhereJoinType][]*Wheres
	wheres    map[contracts.WhereJoinType][]*Where
}

func (this *Wheres) IsEmpty() bool {
	return len(this.subWheres) == 0 && len(this.wheres) == 0
}

func (this Wheres) getSubStringers(whereType contracts.WhereJoinType) []fmt.Stringer {
	stringers := make([]fmt.Stringer, 0)
	for _, where := range this.subWheres[whereType] {
		stringers = append(stringers, where)
	}
	return stringers
}

func (this Wheres) getStringers(whereType contracts.WhereJoinType) []fmt.Stringer {
	stringers := make([]fmt.Stringer, 0)
	for _, where := range this.wheres[whereType] {
		stringers = append(stringers, where)
	}
	return stringers
}

func (this *Wheres) getSubWheres(whereType contracts.WhereJoinType) string {
	return JoinSubStringerArray(this.getSubStringers(whereType), string(whereType))
}

func (this *Wheres) getWheres(whereType contracts.WhereJoinType) string {
	return JoinStringerArray(this.getStringers(whereType), string(whereType))
}

func (this *Wheres) String() (result string) {
	if this == nil || this.IsEmpty() {
		return ""
	}

	result = this.getSubWheres(contracts.And)
	andWheres := this.getWheres(contracts.And)

	if result != "" {
		if andWheres != "" {
			result = fmt.Sprintf("%s and %s", result, andWheres)
		}
	} else {
		result = andWheres
	}

	orSubWheres := this.getSubWheres(contracts.Or)
	if result == "" {
		result = orSubWheres
	} else if orSubWheres != "" {
		result = fmt.Sprintf("%s or %s", result, orSubWheres)
	}

	orWheres := this.getWheres(contracts.Or)
	if result == "" {
		result = orWheres
	} else if orWheres != "" {
		result = fmt.Sprintf("%s or %s", result, orWheres)
	}

	return
}

func (builder *Builder) WhereFunc(callback contracts.QueryFunc, whereType ...contracts.WhereJoinType) contracts.QueryBuilder {
	subBuilder := &Builder{
		wheres: &Wheres{
			wheres:    map[contracts.WhereJoinType][]*Where{},
			subWheres: map[contracts.WhereJoinType][]*Wheres{},
		},
		bindings: map[bindingType][]interface{}{},
	}
	callback(subBuilder)
	if len(whereType) == 0 {
		builder.wheres.subWheres[contracts.And] = append(builder.wheres.subWheres[contracts.And], subBuilder.getWheres())
	} else {
		builder.wheres.subWheres[whereType[0]] = append(builder.wheres.subWheres[whereType[0]], subBuilder.getWheres())
	}
	return builder.addBinding(whereBinding, subBuilder.GetBindings()...)
}

func (builder *Builder) WhereFields(fields contracts.Fields) contracts.QueryBuilder {
	for column, value := range fields {
		builder.Where(column, value)
	}
	return builder
}

func (builder *Builder) OrWhereFunc(callback contracts.QueryFunc) contracts.QueryBuilder {
	return builder.WhereFunc(callback, contracts.Or)
}

func (builder *Builder) Where(field string, args ...interface{}) contracts.QueryBuilder {
	var (
		arg       interface{}
		condition = "="
		whereType = contracts.And
	)
	switch len(args) {
	case 1:
		arg = args[0]
	case 2:
		condition = args[0].(string)
		arg = args[1]
	case 3:
		condition = args[0].(string)
		arg = args[1]
		whereType = args[2].(contracts.WhereJoinType)
	}

	raw, bindings := builder.prepareArgs(condition, arg)

	builder.wheres.wheres[whereType] = append(builder.wheres.wheres[whereType], &Where{
		field:     field,
		condition: condition,
		arg:       raw,
	})

	return builder.addBinding(whereBinding, bindings...)
}

func (builder *Builder) OrWhere(field string, args ...interface{}) contracts.QueryBuilder {
	var (
		arg       interface{}
		condition = "="
	)
	switch len(args) {
	case 1:
		arg = args[0]
	case 2:
		condition = args[0].(string)
		arg = args[1]
	default:
		condition = args[0].(string)
		arg = args[1]
	}
	raw, bindings := builder.prepareArgs(condition, arg)

	builder.wheres.wheres[contracts.Or] = append(builder.wheres.wheres[contracts.Or], &Where{
		field:     field,
		condition: condition,
		arg:       raw,
	})
	return builder.addBinding(whereBinding, bindings...)
}

func JoinStringerArray(arr []fmt.Stringer, sep string) (result string) {
	for index, stringer := range arr {
		if index == 0 {
			result = stringer.String()
		} else {
			result = fmt.Sprintf("%s %s %s", result, sep, stringer.String())
		}
	}

	return
}

func JoinSubStringerArray(arr []fmt.Stringer, sep string) (result string) {
	for index, stringer := range arr {
		if index == 0 {
			result = fmt.Sprintf("(%s)", stringer.String())
		} else {
			result = fmt.Sprintf("%s %s (%s)", result, sep, stringer.String())
		}
	}

	return
}
