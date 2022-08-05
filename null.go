package querybuilder

import "github.com/goal-web/contracts"

func (this *Builder[T]) WhereIsNull(field string, whereType ...contracts.WhereJoinType) contracts.QueryBuilder[T] {
	if len(whereType) == 0 {
		return this.Where(field, "is", "null")
	}
	return this.Where(field, "is", "null", whereType[0])
}

func (this *Builder[T]) WhereNotNull(field string, whereType ...contracts.WhereJoinType) contracts.QueryBuilder[T] {
	if len(whereType) == 0 {
		return this.Where(field, "is not", "null")
	}
	return this.Where(field, "is not", "null", whereType[0])
}

func (this *Builder[T]) OrWhereIsNull(field string) contracts.QueryBuilder[T] {
	return this.OrWhere(field, "is", "null")
}

func (this *Builder[T]) OrWhereNotNull(field string) contracts.QueryBuilder[T] {
	return this.OrWhere(field, "is not", "null")
}
