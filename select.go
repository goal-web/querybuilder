package querybuilder

import (
	"fmt"
	"github.com/goal-web/contracts"
)

func (builder *Builder[T]) Select(fields ...string) contracts.Query[T] {
	builder.fields = fields
	return builder
}

func (builder *Builder[T]) AddSelect(fields ...string) contracts.Query[T] {
	builder.fields = append(builder.fields, fields...)
	return builder
}

func (builder *Builder[T]) SelectSub(provider contracts.QueryProvider[T], as string) contracts.Query[T] {
	subBuilder := provider()
	builder.fields = []string{fmt.Sprintf("(%s) as %s", subBuilder.ToSql(), as)}
	return builder.addBinding(selectBinding, subBuilder.GetBindings()...)
}

func (builder *Builder[T]) AddSelectSub(provider contracts.QueryProvider[T], as string) contracts.Query[T] {
	subBuilder := provider()
	builder.fields = append(builder.fields, fmt.Sprintf("(%s) as %s", subBuilder.ToSql(), as))
	return builder.addBinding(selectBinding, subBuilder.GetBindings()...)
}
