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

func (this *Builder) Union(builder contracts.QueryBuilder, unionType ...contracts.UnionJoinType) contracts.QueryBuilder {
	if builder != nil {
		if len(unionType) > 0 {
			this.unions[unionType[0]] = append(this.unions[unionType[0]], builder)
		} else {
			this.unions[contracts.Union] = append(this.unions[contracts.Union], builder)
		}
	}

	return this.addBinding(unionBinding, builder.GetBindings()...)
}

func (this *Builder) UnionAll(builder contracts.QueryBuilder) contracts.QueryBuilder {
	return this.Union(builder, contracts.UnionAll)
}

func (this *Builder) UnionByProvider(builder contracts.QueryProvider, unionType ...contracts.UnionJoinType) contracts.QueryBuilder {
	return this.Union(builder(), unionType...)
}

func (this *Builder) UnionAllByProvider(builder contracts.QueryProvider) contracts.QueryBuilder {
	return this.Union(builder(), contracts.UnionAll)
}
