package querybuilder

import (
	"fmt"
	"github.com/goal-web/contracts"
)

func NewBuilder[T any](table string) *Builder[T] {
	return &Builder[T]{
		table:    table,
		fields:   []string{"*"},
		orderBy:  OrderByFields{},
		bindings: map[bindingType][]any{},
		joins:    Joins{},
		unions:   Unions[T]{},
		groupBy:  GroupBy{},
		wheres: &Wheres{
			wheres:    map[contracts.WhereJoinType][]*Where{},
			subWheres: map[contracts.WhereJoinType][]*Wheres{},
		},
		having: &Wheres{
			wheres:    map[contracts.WhereJoinType][]*Where{},
			subWheres: map[contracts.WhereJoinType][]*Wheres{},
		},
	}
}

func New[T any](table string) contracts.Query[T] {
	return &Builder[T]{
		table:    table,
		fields:   []string{"*"},
		orderBy:  OrderByFields{},
		bindings: map[bindingType][]any{},
		joins:    Joins{},
		unions:   Unions[T]{},
		groupBy:  GroupBy{},
		wheres: &Wheres{
			wheres:    map[contracts.WhereJoinType][]*Where{},
			subWheres: map[contracts.WhereJoinType][]*Wheres{},
		},
		having: &Wheres{
			wheres:    map[contracts.WhereJoinType][]*Where{},
			subWheres: map[contracts.WhereJoinType][]*Wheres{},
		},
	}
}

func FromSub[T any](callback contracts.QueryProvider[T], as string) contracts.Query[T] {
	subQuery := callback()
	builder := NewBuilder[T]("")
	builder.table = fmt.Sprintf("(%s) as %s", subQuery.ToSql(), as)
	return builder.addBinding(fromBinding, subQuery.GetBindings()...)
}

func FromQuery[T any](subQuery contracts.QueryBuilder[T], as string) contracts.Query[T] {
	builder := NewBuilder[T]("")
	builder.table = fmt.Sprintf("(%s) as %s", subQuery.ToSql(), as)
	return builder.addBinding(fromBinding, subQuery.GetBindings()...)
}
