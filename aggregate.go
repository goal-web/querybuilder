package querybuilder

import (
	"fmt"
	"github.com/goal-web/contracts"
)

func (this *Builder) WithCount(fields ...string) contracts.QueryBuilder {
	if len(fields) == 0 {
		return this.Select("count(*)")
	}
	return this.Select(fmt.Sprintf("count(%s) as %s_count", fields[0], fields[0]))
}

func (this *Builder) WithAvg(field string, as ...string) contracts.QueryBuilder {
	if len(as) == 0 {
		return this.Select(fmt.Sprintf("avg(%s) as %s_avg", field, field))
	}
	return this.Select(fmt.Sprintf("avg(%s) as %s", field, as[0]))
}

func (this *Builder) WithSum(field string, as ...string) contracts.QueryBuilder {
	if len(as) == 0 {
		return this.Select(fmt.Sprintf("sum(%s) as %s_sum", field, field))
	}
	return this.Select(fmt.Sprintf("sum(%s) as %s", field, as[0]))
}

func (this *Builder) WithMax(field string, as ...string) contracts.QueryBuilder {
	if len(as) == 0 {
		return this.Select(fmt.Sprintf("max(%s) as %s_max", field, field))
	}
	return this.Select(fmt.Sprintf("max(%s) as %s", field, as[0]))
}

func (this *Builder) WithMin(field string, as ...string) contracts.QueryBuilder {
	if len(as) == 0 {
		return this.Select(fmt.Sprintf("min(%s) as %s_min", field, field))
	}
	return this.Select(fmt.Sprintf("min(%s) as %s", field, as[0]))
}
