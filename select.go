package querybuilder

import (
	"fmt"
	"github.com/goal-web/contracts"
)

func (this *Builder[T]) Select(fields ...string) contracts.QueryBuilder[T] {
	this.fields = fields
	return this
}

func (this *Builder[T]) AddSelect(fields ...string) contracts.QueryBuilder[T] {
	this.fields = append(this.fields, fields...)
	return this
}

func (this *Builder[T]) SelectSub(provider contracts.QueryProvider[T], as string) contracts.QueryBuilder[T] {
	subBuilder := provider()
	this.fields = []string{fmt.Sprintf("(%s) as %s", subBuilder.ToSql(), as)}
	return this.addBinding(selectBinding, subBuilder.GetBindings()...)
}
func (this *Builder[T]) AddSelectSub(provider contracts.QueryProvider[T], as string) contracts.QueryBuilder[T] {
	subBuilder := provider()
	this.fields = append(this.fields, fmt.Sprintf("(%s) as %s", subBuilder.ToSql(), as))
	return this.addBinding(selectBinding, subBuilder.GetBindings()...)
}
