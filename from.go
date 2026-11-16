package querybuilder

import "github.com/goal-web/contracts"

type fromSub[T any] struct {
	as       string
	subQuery contracts.QueryBuilder[T]
}
