package querybuilder

import "github.com/goal-web/contracts"

func (this *Builder) WhereBetween(field string, args interface{}, whereType ...contracts.WhereJoinType) contracts.QueryBuilder {
	if len(whereType) > 0 {
		return this.Where(field, "between", args, whereType[0])
	}

	return this.Where(field, "between", args)
}

func (this *Builder) OrWhereBetween(field string, args interface{}) contracts.QueryBuilder {
	return this.OrWhere(field, "between", args)
}

func (this *Builder) WhereNotBetween(field string, args interface{}, whereType ...contracts.WhereJoinType) contracts.QueryBuilder {
	if len(whereType) > 0 {
		return this.Where(field, "not between", args, whereType[0])
	}

	return this.Where(field, "not between", args)
}

func (this *Builder) OrWhereNotBetween(field string, args interface{}) contracts.QueryBuilder {
	return this.OrWhere(field, "not between", args)
}
