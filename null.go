package querybuilder

import "github.com/goal-web/contracts"

func (builder *Builder[T]) WhereIsNull(field string, whereType ...contracts.WhereJoinType) contracts.Query[T] {
	if len(whereType) == 0 {
		return builder.Where(field, "is", "null")
	}
	return builder.Where(field, "is", "null", whereType[0])
}

func (builder *Builder[T]) WhereNotNull(field string, whereType ...contracts.WhereJoinType) contracts.Query[T] {
	if len(whereType) == 0 {
		return builder.Where(field, "is not", "null")
	}
	return builder.Where(field, "is not", "null", whereType[0])
}

func (builder *Builder[T]) OrWhereIsNull(field string) contracts.Query[T] {
	return builder.OrWhere(field, "is", "null")
}

func (builder *Builder[T]) OrWhereNotNull(field string) contracts.Query[T] {
	return builder.OrWhere(field, "is not", "null")
}
