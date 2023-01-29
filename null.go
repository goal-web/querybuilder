package querybuilder

import "github.com/goal-web/contracts"

func (builder *Builder) WhereIsNull(field string, whereType ...contracts.WhereJoinType) contracts.QueryBuilder {
	if len(whereType) == 0 {
		return builder.Where(field, "is", "null")
	}
	return builder.Where(field, "is", "null", whereType[0])
}

func (builder *Builder) WhereNotNull(field string, whereType ...contracts.WhereJoinType) contracts.QueryBuilder {
	if len(whereType) == 0 {
		return builder.Where(field, "is not", "null")
	}
	return builder.Where(field, "is not", "null", whereType[0])
}

func (builder *Builder) OrWhereIsNull(field string) contracts.QueryBuilder {
	return builder.OrWhere(field, "is", "null")
}

func (builder *Builder) OrWhereNotNull(field string) contracts.QueryBuilder {
	return builder.OrWhere(field, "is not", "null")
}
