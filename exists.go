package querybuilder

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
)

func (builder *Builder[T]) WhereExists(provider contracts.QueryProvider[T], where ...contracts.WhereJoinType) contracts.QueryBuilder[T] {
	subBuilder := provider()
	subSql := fmt.Sprintf("(%s)", subBuilder.ToSql())

	return builder.addBinding(whereBinding, subBuilder.GetBindings()...).
		Where("", "exists", subSql, utils.DefaultValue(where, contracts.And))
}

func (builder *Builder[T]) WhereExistsRaw(raw string, where ...contracts.WhereJoinType) contracts.QueryBuilder[T] {
	return builder.Where("", "exists", raw, utils.DefaultValue(where, contracts.And))
}

func (builder *Builder[T]) OrWhereExistsRaw(raw string) contracts.QueryBuilder[T] {
	return builder.Where("", "exists", raw, contracts.Or)
}

func (builder *Builder[T]) OrWhereExists(provider contracts.QueryProvider[T]) contracts.QueryBuilder[T] {
	return builder.WhereExists(provider, contracts.Or)
}

func (builder *Builder[T]) WhereNotExists(provider contracts.QueryProvider[T], where ...contracts.WhereJoinType) contracts.QueryBuilder[T] {
	subBuilder := provider()
	subSql := fmt.Sprintf("(%s)", subBuilder.ToSql())
	if len(where) == 0 {
		return builder.addBinding(whereBinding, subBuilder.GetBindings()...).
			Where("", "not exists", subSql)
	}

	return builder.addBinding(whereBinding, subBuilder.GetBindings()...).
		Where("", "not exists", subSql, where[0])
}

func (builder *Builder[T]) OrWhereNotExists(provider contracts.QueryProvider[T]) contracts.QueryBuilder[T] {
	return builder.WhereNotExists(provider, contracts.Or)
}
