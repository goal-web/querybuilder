package querybuilder

import (
	"fmt"
	"github.com/goal-web/contracts"
)

func (builder *Builder) WithCount(fields ...string) contracts.QueryBuilder {
	if len(fields) == 0 {
		return builder.Select("count(*)")
	}
	return builder.Select(fmt.Sprintf("count(%s) as %s_count", fields[0], fields[0]))
}

func (builder *Builder) WithAvg(field string, as ...string) contracts.QueryBuilder {
	if len(as) == 0 {
		return builder.Select(fmt.Sprintf("avg(%s) as %s_avg", field, field))
	}
	return builder.Select(fmt.Sprintf("avg(%s) as %s", field, as[0]))
}

func (builder *Builder) WithSum(field string, as ...string) contracts.QueryBuilder {
	if len(as) == 0 {
		return builder.Select(fmt.Sprintf("sum(%s) as %s_sum", field, field))
	}
	return builder.Select(fmt.Sprintf("sum(%s) as %s", field, as[0]))
}

func (builder *Builder) WithMax(field string, as ...string) contracts.QueryBuilder {
	if len(as) == 0 {
		return builder.Select(fmt.Sprintf("max(%s) as %s_max", field, field))
	}
	return builder.Select(fmt.Sprintf("max(%s) as %s", field, as[0]))
}

func (builder *Builder) WithMin(field string, as ...string) contracts.QueryBuilder {
	if len(as) == 0 {
		return builder.Select(fmt.Sprintf("min(%s) as %s_min", field, field))
	}
	return builder.Select(fmt.Sprintf("min(%s) as %s", field, as[0]))
}
