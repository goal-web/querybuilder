package querybuilder

import (
	"fmt"
	"github.com/goal-web/contracts"
)

func (this *Builder[T]) WhereExists(provider contracts.QueryProvider[T], where ...contracts.WhereJoinType) contracts.QueryBuilder[T] {
	subBuilder := provider()
	subSql := fmt.Sprintf("(%s)", subBuilder.ToSql())
	if len(where) == 0 {
		return this.addBinding(whereBinding, subBuilder.GetBindings()...).
			Where("", "exists", subSql)
	}

	return this.addBinding(whereBinding, subBuilder.GetBindings()...).
		Where("", "exists", subSql, where[0])

}

func (this *Builder[T]) OrWhereExists(provider contracts.QueryProvider[T]) contracts.QueryBuilder[T] {
	return this.WhereExists(provider, contracts.Or)
}

func (this *Builder[T]) WhereNotExists(provider contracts.QueryProvider[T], where ...contracts.WhereJoinType) contracts.QueryBuilder[T] {
	subBuilder := provider()
	subSql := fmt.Sprintf("(%s)", subBuilder.ToSql())
	if len(where) == 0 {
		return this.addBinding(whereBinding, subBuilder.GetBindings()...).
			Where("", "not exists", subSql)
	}

	return this.addBinding(whereBinding, subBuilder.GetBindings()...).
		Where("", "not exists", subSql, where[0])
}

func (this *Builder[T]) OrWhereNotExists(provider contracts.QueryProvider[T]) contracts.QueryBuilder[T] {
	return this.WhereNotExists(provider, contracts.Or)
}
