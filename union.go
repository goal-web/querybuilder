package querybuilder

import (
	"fmt"
	"github.com/goal-web/contracts"
)

type Unions[T any] map[contracts.UnionJoinType][]contracts.QueryBuilder[T]

func (unions Unions[T]) IsEmpty() bool {
	return len(unions) == 0
}

func (unions Unions[T]) String() (result string) {
	if unions.IsEmpty() {
		return
	}
	for unionType, builders := range unions {
		for _, builder := range builders {
			result = fmt.Sprintf("%s %s (%s)", result, unionType, builder.ToSql())
		}
	}

	return
}

func (builder *Builder[T]) Union(b contracts.QueryBuilder[T], unionType ...contracts.UnionJoinType) contracts.Query[T] {
	if builder != nil {
		if len(unionType) > 0 {
			builder.unions[unionType[0]] = append(builder.unions[unionType[0]], b)
		} else {
			builder.unions[contracts.Union] = append(builder.unions[contracts.Union], b)
		}
	}

	return builder.addBinding(unionBinding, builder.GetBindings()...)
}

func (builder *Builder[T]) UnionAll(b contracts.QueryBuilder[T]) contracts.Query[T] {
	return builder.Union(b, contracts.UnionAll)
}

func (builder *Builder[T]) UnionByProvider(provider contracts.QueryProvider[T], unionType ...contracts.UnionJoinType) contracts.Query[T] {
	return builder.Union(provider(), unionType...)
}

func (builder *Builder[T]) UnionAllByProvider(provider contracts.QueryProvider[T]) contracts.Query[T] {
	return builder.Union(provider(), contracts.UnionAll)
}
