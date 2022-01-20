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

func (this *Builder) Join(table string, first, condition, second string, joins ...contracts.JoinType) contracts.QueryBuilder {
	join := contracts.InnerJoin
	if len(joins) > 0 {
		join = joins[0]
	}
	this.joins = append(this.joins, Join{table, join, &Wheres{wheres: map[contracts.WhereJoinType][]*Where{
		contracts.And: {&Where{
			field:     first,
			condition: condition,
			arg:       second,
		}},
	}}})

	return this
}

func (this *Builder) JoinSub(provider contracts.QueryProvider, as, first, condition, second string, joins ...contracts.JoinType) contracts.QueryBuilder {
	join := contracts.InnerJoin
	if len(joins) > 0 {
		join = joins[0]
	}
	subBuilder := provider()
	this.joins = append(this.joins, Join{fmt.Sprintf("(%s) as %s", subBuilder.ToSql(), as), join, &Wheres{wheres: map[contracts.WhereJoinType][]*Where{
		contracts.And: {&Where{
			field:     first,
			condition: condition,
			arg:       second,
		}},
	}}})

	return this.addBinding(joinBinding, subBuilder.GetBindings()...)
}

func (this *Builder) FullJoin(table string, first, condition, second string) contracts.QueryBuilder {
	return this.Join(table, first, condition, second, contracts.FullJoin)
}
func (this *Builder) FullOutJoin(table string, first, condition, second string) contracts.QueryBuilder {
	return this.Join(table, first, condition, second, contracts.FullOutJoin)
}

func (this *Builder) LeftJoin(table string, first, condition, second string) contracts.QueryBuilder {
	return this.Join(table, first, condition, second, contracts.LeftJoin)
}

func (this *Builder) RightJoin(table string, first, condition, second string) contracts.QueryBuilder {
	return this.Join(table, first, condition, second, contracts.RightJoin)
}
