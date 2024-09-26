package querybuilder

import "github.com/goal-web/contracts"

// WhereBetween args 参数可以是整数、浮点数、字符串、any 等类型的数组，或者用` and `隔开的字符串，或者在源码中了解更多 https://github.com/goal-web/querybuilder/blob/78bcc832604bfcdb68579e3dd1441796a16994cf/builder.go#L74
func (builder *Builder[T]) WhereBetween(field string, args any, whereType ...contracts.WhereJoinType) contracts.QueryBuilder[T] {
	if len(whereType) > 0 {
		return builder.Where(field, "between", args, whereType[0])
	}

	return builder.Where(field, "between", args)
}

func (builder *Builder[T]) OrWhereBetween(field string, args any) contracts.QueryBuilder[T] {
	return builder.OrWhere(field, "between", args)
}

func (builder *Builder[T]) WhereNotBetween(field string, args any, whereType ...contracts.WhereJoinType) contracts.QueryBuilder[T] {
	if len(whereType) > 0 {
		return builder.Where(field, "not between", args, whereType[0])
	}

	return builder.Where(field, "not between", args)
}

func (builder *Builder[T]) OrWhereNotBetween(field string, args any) contracts.QueryBuilder[T] {
	return builder.OrWhere(field, "not between", args)
}
