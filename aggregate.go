package querybuilder

import (
	"fmt"
	"github.com/goal-web/contracts"
)

func (this *Builder[T]) WithCount(fields ...string) contracts.QueryBuilder[T] {
	if len(fields) == 0 {
		return this.Select("count(*)")
	}
	return this.Select(fmt.Sprintf("count(%s) as %s_count", fields[0], fields[0]))
}

func (this *Builder[T]) WithAvg(field string, as ...string) contracts.QueryBuilder[T] {
	if len(as) == 0 {
		return this.Select(fmt.Sprintf("avg(%s) as %s_avg", field, field))
	}
	return this.Select(fmt.Sprintf("avg(%s) as %s", field, as[0]))
}

func (this *Builder[T]) WithSum(field string, as ...string) contracts.QueryBuilder[T] {
	if len(as) == 0 {
		return this.Select(fmt.Sprintf("sum(%s) as %s_sum", field, field))
	}
	return this.Select(fmt.Sprintf("sum(%s) as %s", field, as[0]))
}

func (this *Builder[T]) WithMax(field string, as ...string) contracts.QueryBuilder[T] {
	if len(as) == 0 {
		return this.Select(fmt.Sprintf("max(%s) as %s_max", field, field))
	}
	return this.Select(fmt.Sprintf("max(%s) as %s", field, as[0]))
}

func (this *Builder[T]) WithMin(field string, as ...string) contracts.QueryBuilder[T] {
	if len(as) == 0 {
		return this.Select(fmt.Sprintf("min(%s) as %s_min", field, field))
	}
	return this.Select(fmt.Sprintf("min(%s) as %s", field, as[0]))
}
