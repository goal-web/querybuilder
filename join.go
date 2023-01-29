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

func (this Join) String() (result string) {
	result = fmt.Sprintf("%s join %s", this.join, this.table)
	if this.conditions.IsEmpty() {
		return
	}
	result = fmt.Sprintf("%s on (%s)", result, this.conditions.String())
	return
}

type Joins []Join

func (this Joins) IsEmpty() bool {
	return len(this) == 0
}

func (this Joins) String() (result string) {
	if this.IsEmpty() {
		return
	}

	for index, join := range this {
		if index == 0 {
			result = join.String()
		} else {
			result = fmt.Sprintf("%s %s", result, join.String())
		}
	}

	return
}

func (builder *Builder) Join(table string, first, condition, second string, joins ...contracts.JoinType) contracts.QueryBuilder {
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

func (builder *Builder) JoinSub(provider contracts.QueryProvider, as, first, condition, second string, joins ...contracts.JoinType) contracts.QueryBuilder {
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

func (builder *Builder) FullJoin(table string, first, condition, second string) contracts.QueryBuilder {
	return builder.Join(table, first, condition, second, contracts.FullJoin)
}
func (builder *Builder) FullOutJoin(table string, first, condition, second string) contracts.QueryBuilder {
	return builder.Join(table, first, condition, second, contracts.FullOutJoin)
}

func (builder *Builder) LeftJoin(table string, first, condition, second string) contracts.QueryBuilder {
	return builder.Join(table, first, condition, second, contracts.LeftJoin)
}

func (builder *Builder) RightJoin(table string, first, condition, second string) contracts.QueryBuilder {
	return builder.Join(table, first, condition, second, contracts.RightJoin)
}
