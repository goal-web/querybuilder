package querybuilder

import (
	"fmt"
	"github.com/goal-web/contracts"
)

func (builder *Builder) Select(fields ...string) contracts.QueryBuilder {
	builder.fields = fields
	return builder
}

func (builder *Builder) AddSelect(fields ...string) contracts.QueryBuilder {
	builder.fields = append(builder.fields, fields...)
	return builder
}

func (builder *Builder) SelectSub(provider contracts.QueryProvider, as string) contracts.QueryBuilder {
	subBuilder := provider()
	builder.fields = []string{fmt.Sprintf("(%s) as %s", subBuilder.ToSql(), as)}
	return builder.addBinding(selectBinding, subBuilder.GetBindings()...)
}
func (builder *Builder) AddSelectSub(provider contracts.QueryProvider, as string) contracts.QueryBuilder {
	subBuilder := provider()
	builder.fields = append(builder.fields, fmt.Sprintf("(%s) as %s", subBuilder.ToSql(), as))
	return builder.addBinding(selectBinding, subBuilder.GetBindings()...)
}
