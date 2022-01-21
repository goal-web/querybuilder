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
func (this *Builder) SimplePaginate(perPage int64, current ...int64) interface{} {
	return this.WithPagination(perPage, current...).Get()
}

func (this *Builder) FirstOr(provider contracts.InstanceProvider) interface{} {
	if result := this.First(); result != nil {
		return result
	}
	return provider()
}

func (this *Builder) FirstWhere(column string, args ...interface{}) interface{} {
	return this.Where(column, args...).First()
}

func (this *Builder) Paginate(perPage int64, current ...int64) (interface{}, int64) {
	return this.SimplePaginate(perPage, current...), this.Count()
}
