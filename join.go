package querybuilder

import (
	"fmt"
	"github.com/goal-web/contracts"
)

type Join struct {
	table      string
	join       contracts.JoinType
	conditions *Wheres
}

func (join Join) String() (result string) {
	result = fmt.Sprintf("%s join %s", join.join, join.table)
	if join.conditions.IsEmpty() {
		return
	}
	result = fmt.Sprintf("%s on (%s)", result, join.conditions.String())
	return
}

type Joins []Join

func (joins Joins) IsEmpty() bool {
	return len(joins) == 0
}

func (joins Joins) String() (result string) {
	if joins.IsEmpty() {
		return
	}

	for index, join := range joins {
		if index == 0 {
			result = join.String()
		} else {
			result = fmt.Sprintf("%s %s", result, join.String())
		}
	}

	return
}

func (builder *Builder[T]) Join(table string, first, condition, second string, joins ...contracts.JoinType) contracts.Query[T] {
	join := contracts.InnerJoin
	if len(joins) > 0 {
		join = joins[0]
	}
	builder.joins = append(builder.joins, Join{table, join, &Wheres{wheres: map[contracts.WhereJoinType][]*Where{
		contracts.And: {&Where{
			field:     first,
			condition: condition,
			arg:       second,
		}},
	}}})

	return builder
}

func (builder *Builder[T]) JoinSub(provider contracts.QueryProvider[T], as, first, condition, second string, joins ...contracts.JoinType) contracts.Query[T] {
	join := contracts.InnerJoin
	if len(joins) > 0 {
		join = joins[0]
	}
	subBuilder := provider()
	builder.joins = append(builder.joins, Join{fmt.Sprintf("(%s) as %s", subBuilder.ToSql(), as), join, &Wheres{wheres: map[contracts.WhereJoinType][]*Where{
		contracts.And: {&Where{
			field:     first,
			condition: condition,
			arg:       second,
		}},
	}}})

	return builder.addBinding(joinBinding, subBuilder.GetBindings()...)
}

func (builder *Builder[T]) FullJoin(table string, first, condition, second string) contracts.Query[T] {
	return builder.Join(table, first, condition, second, contracts.FullJoin)
}
func (builder *Builder[T]) FullOutJoin(table string, first, condition, second string) contracts.Query[T] {
	return builder.Join(table, first, condition, second, contracts.FullOutJoin)
}

func (builder *Builder[T]) LeftJoin(table string, first, condition, second string) contracts.Query[T] {
	return builder.Join(table, first, condition, second, contracts.LeftJoin)
}

func (builder *Builder[T]) RightJoin(table string, first, condition, second string) contracts.Query[T] {
	return builder.Join(table, first, condition, second, contracts.RightJoin)
}
