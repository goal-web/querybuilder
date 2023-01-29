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

func (builder *Builder) GroupBy(columns ...string) contracts.QueryBuilder {
	builder.groupBy = append(builder.groupBy, columns...)

	return builder
}

func (builder *Builder) Having(field string, args ...interface{}) contracts.QueryBuilder {
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

	builder.having.wheres[whereType] = append(builder.having.wheres[whereType], &Where{
		field:     field,
		condition: condition,
		arg:       raw,
	})

	return builder.addBinding(havingBinding, bindings...)
}

func (builder *Builder) OrHaving(field string, args ...interface{}) contracts.QueryBuilder {
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

	builder.having.wheres[contracts.Or] = append(builder.having.wheres[contracts.Or], &Where{
		field:     field,
		condition: condition,
		arg:       raw,
	})
	return builder.addBinding(havingBinding, bindings...)
}
