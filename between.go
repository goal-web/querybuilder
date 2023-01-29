package querybuilder

import "github.com/goal-web/contracts"

// WhereBetween args 参数可以是整数、浮点数、字符串、interface{} 等类型的数组，或者用` and `隔开的字符串，或者在源码中了解更多 https://github.com/goal-web/querybuilder/blob/78bcc832604bfcdb68579e3dd1441796a16994cf/builder.go#L74
func (builder *Builder) WhereBetween(field string, args interface{}, whereType ...contracts.WhereJoinType) contracts.QueryBuilder {
	if len(whereType) > 0 {
		return builder.Where(field, "between", args, whereType[0])
	}

	return builder.Where(field, "between", args)
}

func (builder *Builder) OrWhereBetween(field string, args interface{}) contracts.QueryBuilder {
	return builder.OrWhere(field, "between", args)
}

func (builder *Builder) WhereNotBetween(field string, args interface{}, whereType ...contracts.WhereJoinType) contracts.QueryBuilder {
	if len(whereType) > 0 {
		return builder.Where(field, "not between", args, whereType[0])
	}

	return builder.Where(field, "not between", args)
}

func (builder *Builder) OrWhereNotBetween(field string, args interface{}) contracts.QueryBuilder {
	return builder.OrWhere(field, "not between", args)
}
