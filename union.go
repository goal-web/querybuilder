package querybuilder

import (
	"fmt"
	"github.com/goal-web/contracts"
)

type Unions map[contracts.UnionJoinType][]contracts.QueryBuilder

func (this Unions) IsEmpty() bool {
	return len(this) == 0
}

func (this Unions) String() (result string) {
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

func (builder *Builder) Union(b contracts.QueryBuilder, unionType ...contracts.UnionJoinType) contracts.QueryBuilder {
	if builder != nil {
		if len(unionType) > 0 {
			builder.unions[unionType[0]] = append(builder.unions[unionType[0]], b)
		} else {
			builder.unions[contracts.Union] = append(builder.unions[contracts.Union], b)
		}
	}

	return builder.addBinding(unionBinding, builder.GetBindings()...)
}

func (builder *Builder) UnionAll(b contracts.QueryBuilder) contracts.QueryBuilder {
	return builder.Union(b, contracts.UnionAll)
}

func (builder *Builder) UnionByProvider(provider contracts.QueryProvider, unionType ...contracts.UnionJoinType) contracts.QueryBuilder {
	return builder.Union(provider(), unionType...)
}

func (builder *Builder) UnionAllByProvider(provider contracts.QueryProvider) contracts.QueryBuilder {
	return builder.Union(provider(), contracts.UnionAll)
}
