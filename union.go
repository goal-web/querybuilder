package querybuilder

import (
	"fmt"
	"github.com/goal-web/contracts"
)

type Unions[T contracts.GetFields] map[contracts.UnionJoinType][]contracts.QueryBuilder[T]

func (this Unions[T]) IsEmpty() bool {
	return len(this) == 0
}

func (this Unions[T]) String() (result string) {
	if this.IsEmpty() {
		return
	}
	for unionType, builders := range this {
		for _, builder := range builders {
			result = fmt.Sprintf("%s %s (%s)", result, unionType, builder.ToSql())
		}
	}

	return
}

func (this *Builder[T]) Union(builder contracts.QueryBuilder[T], unionType ...contracts.UnionJoinType) contracts.QueryBuilder[T] {
	if builder != nil {
		if len(unionType) > 0 {
			this.unions[unionType[0]] = append(this.unions[unionType[0]], builder)
		} else {
			this.unions[contracts.Union] = append(this.unions[contracts.Union], builder)
		}
	}

	return this.addBinding(unionBinding, builder.GetBindings()...)
}

func (this *Builder[T]) UnionAll(builder contracts.QueryBuilder[T]) contracts.QueryBuilder[T] {
	return this.Union(builder, contracts.UnionAll)
}

func (this *Builder[T]) UnionByProvider(builder contracts.QueryProvider[T], unionType ...contracts.UnionJoinType) contracts.QueryBuilder[T] {
	return this.Union(builder(), unionType...)
}

func (this *Builder[T]) UnionAllByProvider(builder contracts.QueryProvider[T]) contracts.QueryBuilder[T] {
	return this.Union(builder(), contracts.UnionAll)
}
