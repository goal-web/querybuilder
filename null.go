package querybuilder

import "github.com/goal-web/contracts"

func (this *Builder) WhereIsNull(field string, whereType ...contracts.WhereJoinType) contracts.QueryBuilder {
	if len(whereType) == 0 {
		return this.Where(field, "is", "null")
	}
	return this.Where(field, "is", "null", whereType[0])
}

func (this *Builder) WhereNotNull(field string, whereType ...contracts.WhereJoinType) contracts.QueryBuilder {
	if len(whereType) == 0 {
		return this.Where(field, "is not", "null")
	}
	return this.Where(field, "is not", "null", whereType[0])
}

func (this *Builder) OrWhereIsNull(field string) contracts.QueryBuilder {
	return this.OrWhere(field, "is", "null")
}

func (this *Builder) OrWhereNotNull(field string) contracts.QueryBuilder {
	return this.OrWhere(field, "is not", "null")
}
