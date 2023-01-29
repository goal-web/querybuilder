package querybuilder

import (
	"fmt"
	"github.com/goal-web/contracts"
)

func (builder *Builder) WhereExists(provider contracts.QueryProvider, where ...contracts.WhereJoinType) contracts.QueryBuilder {
	subBuilder := provider()
	subSql := fmt.Sprintf("(%s)", subBuilder.ToSql())
	if len(where) == 0 {
		return builder.addBinding(whereBinding, subBuilder.GetBindings()...).
			Where("", "exists", subSql)
	}

	return builder.addBinding(whereBinding, subBuilder.GetBindings()...).
		Where("", "exists", subSql, where[0])

}

func (builder *Builder) OrWhereExists(provider contracts.QueryProvider) contracts.QueryBuilder {
	return builder.WhereExists(provider, contracts.Or)
}

func (builder *Builder) WhereNotExists(provider contracts.QueryProvider, where ...contracts.WhereJoinType) contracts.QueryBuilder {
	subBuilder := provider()
	subSql := fmt.Sprintf("(%s)", subBuilder.ToSql())
	if len(where) == 0 {
		return builder.addBinding(whereBinding, subBuilder.GetBindings()...).
			Where("", "not exists", subSql)
	}

	return builder.addBinding(whereBinding, subBuilder.GetBindings()...).
		Where("", "not exists", subSql, where[0])
}

func (builder *Builder) OrWhereNotExists(provider contracts.QueryProvider) contracts.QueryBuilder {
	return builder.WhereNotExists(provider, contracts.Or)
}
