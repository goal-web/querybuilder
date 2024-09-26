package querybuilder

import (
	"fmt"
	"github.com/goal-web/contracts"
)

func (builder *Builder[T]) Select(fields ...string) contracts.QueryBuilder[T] {
	builder.Selects = fields
	return builder
}

func (builder *Builder[T]) AddSelect(fields ...string) contracts.QueryBuilder[T] {
	builder.Selects = append(builder.Selects, fields...)
	return builder
}

func (builder *Builder[T]) SelectSub(provider contracts.QueryProvider[T], as string) contracts.QueryBuilder[T] {
	subBuilder := provider()
	builder.Selects = []string{fmt.Sprintf("(%s) as %s", subBuilder.ToSql(), as)}
	return builder.addBinding(selectBinding, subBuilder.GetBindings()...)
}

func (builder *Builder[T]) AddSelectSub(provider contracts.QueryProvider[T], as string) contracts.QueryBuilder[T] {
	subBuilder := provider()
	builder.Selects = append(builder.Selects, fmt.Sprintf("(%s) as %s", subBuilder.ToSql(), as))
	return builder.addBinding(selectBinding, subBuilder.GetBindings()...)
}
