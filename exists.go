package querybuilder

import (
	"fmt"
	"github.com/goal-web/contracts"
)

func (this *Builder) WhereExists(provider contracts.QueryProvider, where ...contracts.WhereJoinType) contracts.QueryBuilder {
	subBuilder := provider()
	subSql := fmt.Sprintf("(%s)", subBuilder.ToSql())
	if len(where) == 0 {
		return this.addBinding(whereBinding, subBuilder.GetBindings()...).
			Where("", "exists", subSql)
	}

	return this.addBinding(whereBinding, subBuilder.GetBindings()...).
		Where("", "exists", subSql, where[0])

}

func (this *Builder) OrWhereExists(provider contracts.QueryProvider) contracts.QueryBuilder {
	return this.WhereExists(provider, contracts.Or)
}

func (this *Builder) WhereNotExists(provider contracts.QueryProvider, where ...contracts.WhereJoinType) contracts.QueryBuilder {
	subBuilder := provider()
	subSql := fmt.Sprintf("(%s)", subBuilder.ToSql())
	if len(where) == 0 {
		return this.addBinding(whereBinding, subBuilder.GetBindings()...).
			Where("", "not exists", subSql)
	}

	return this.addBinding(whereBinding, subBuilder.GetBindings()...).
		Where("", "not exists", subSql, where[0])
}

func (this *Builder) OrWhereNotExists(provider contracts.QueryProvider) contracts.QueryBuilder {
	return this.WhereNotExists(provider, contracts.Or)
}
