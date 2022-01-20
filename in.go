package querybuilder

import "github.com/goal-web/contracts"

// WhereIn args 参数可以是整数、浮点数、字符串、interface{} 等类型的数组，或者用` and `隔开的字符串，或者在源码中了解更多 https://github.com/goal-web/querybuilder/blob/78bcc832604bfcdb68579e3dd1441796a16994cf/builder.go#L74
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
