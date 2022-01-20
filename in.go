package querybuilder

import "github.com/goal-web/contracts"

func (this *Builder) WhereIn(field string, args interface{}, joinType ...contracts.WhereJoinType) contracts.QueryBuilder {
	if len(joinType) == 0 {
		return this.Where(field, "in", args)
	}
	return this.Where(field, "in", args, joinType[0])
}

func (this *Builder) OrWhereIn(field string, args interface{}) contracts.QueryBuilder {
	return this.OrWhere(field, "in", args)
}

func (this *Builder) WhereNotIn(field string, args interface{}, joinType ...contracts.WhereJoinType) contracts.QueryBuilder {
	if len(joinType) == 0 {
		return this.Where(field, "not in", args)
	}
	return this.Where(field, "not in", args)
}

func (this *Builder) OrWhereNotIn(field string, args interface{}) contracts.QueryBuilder {
	return this.OrWhere(field, "not in", args)
}
