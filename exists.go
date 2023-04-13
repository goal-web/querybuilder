package querybuilder

import (
	"fmt"
	"github.com/goal-web/contracts"
)

func (builder *Builder[T]) WhereExists(provider contracts.QueryProvider[T], where ...contracts.WhereJoinType) contracts.Query[T] {
	subBuilder := provider()
	subSql := fmt.Sprintf("(%s)", subBuilder.ToSql())
	if len(where) == 0 {
		return builder.addBinding(whereBinding, subBuilder.GetBindings()...).
			Where("", "exists", subSql)
	}

	return builder.addBinding(whereBinding, subBuilder.GetBindings()...).
		Where("", "exists", subSql, where[0])

}

func (builder *Builder[T]) OrWhereExists(provider contracts.QueryProvider[T]) contracts.Query[T] {
	return builder.WhereExists(provider, contracts.Or)
}

func (builder *Builder[T]) WhereNotExists(provider contracts.QueryProvider[T], where ...contracts.WhereJoinType) contracts.Query[T] {
	subBuilder := provider()
	subSql := fmt.Sprintf("(%s)", subBuilder.ToSql())
	if len(where) == 0 {
		return builder.addBinding(whereBinding, subBuilder.GetBindings()...).
			Where("", "not exists", subSql)
	}

	return builder.addBinding(whereBinding, subBuilder.GetBindings()...).
		Where("", "not exists", subSql, where[0])
}

func (builder *Builder[T]) OrWhereNotExists(provider contracts.QueryProvider[T]) contracts.Query[T] {
	return builder.WhereNotExists(provider, contracts.Or)
}
