package querybuilder

import (
	"github.com/goal-web/contracts"
	"strings"
)

type GroupBy []string

func (this GroupBy) IsEmpty() bool {
	return len(this) == 0
}

func (this GroupBy) String() string {
	if this.IsEmpty() {
		return ""
	}

	return strings.Join(this, ",")
}

func (this *Builder[T]) GroupBy(columns ...string) contracts.QueryBuilder[T] {
	this.groupBy = append(this.groupBy, columns...)

	return this
}

func (this *Builder[T]) Having(field string, args ...interface{}) contracts.QueryBuilder[T] {
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

	raw, bindings := this.prepareArgs(condition, arg)

	this.having.wheres[whereType] = append(this.having.wheres[whereType], &Where{
		field:     field,
		condition: condition,
		arg:       raw,
	})

	return this.addBinding(havingBinding, bindings...)
}

func (this *Builder[T]) OrHaving(field string, args ...interface{}) contracts.QueryBuilder[T] {
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
	raw, bindings := this.prepareArgs(condition, arg)

	this.having.wheres[contracts.Or] = append(this.having.wheres[contracts.Or], &Where{
		field:     field,
		condition: condition,
		arg:       raw,
	})
	return this.addBinding(havingBinding, bindings...)
}
