package querybuilder

import (
	"fmt"
	"github.com/goal-web/contracts"
)

func (this *Builder) Select(field string, fields ...string) contracts.QueryBuilder {
	this.fields = append(fields, field)
	return this
}

func (this *Builder) AddSelect(fields ...string) contracts.QueryBuilder {
	this.fields = append(this.fields, fields...)
	return this
}

func (this *Builder) SelectSub(provider contracts.QueryProvider, as string) contracts.QueryBuilder {
	subBuilder := provider()
	this.fields = []string{fmt.Sprintf("(%s) as %s", subBuilder.ToSql(), as)}
	return this.addBinding(selectBinding, subBuilder.GetBindings()...)
}
func (this *Builder) AddSelectSub(provider contracts.QueryProvider, as string) contracts.QueryBuilder {
	subBuilder := provider()
	this.fields = append(this.fields, fmt.Sprintf("(%s) as %s", subBuilder.ToSql(), as))
	return this.addBinding(selectBinding, subBuilder.GetBindings()...)
}
