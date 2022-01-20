package querybuilder

import (
	"fmt"
	"github.com/goal-web/contracts"
	"strings"
)

type OrderBy struct {
	field          string
	fieldOrderType contracts.OrderType
}

type OrderByFields []OrderBy

func (this OrderByFields) IsEmpty() bool {
	return len(this) == 0
}

func (this OrderByFields) String() string {
	if this.IsEmpty() {
		return ""
	}

	columns := make([]string, 0)

	for _, orderBy := range this {
		columns = append(columns, fmt.Sprintf("%s %s", orderBy.field, orderBy.fieldOrderType))
	}

	return strings.Join(columns, ",")
}

func (this *Builder) OrderBy(field string, columnOrderType ...contracts.OrderType) contracts.QueryBuilder {
	if len(columnOrderType) > 0 {
		this.orderBy = append(this.orderBy, OrderBy{
			field:          field,
			fieldOrderType: columnOrderType[0],
		})
	} else {
		this.orderBy = append(this.orderBy, OrderBy{
			field:          field,
			fieldOrderType: contracts.Asc,
		})
	}

	return this
}

func (this *Builder) OrderByDesc(field string) contracts.QueryBuilder {
	this.orderBy = append(this.orderBy, OrderBy{
		field:          field,
		fieldOrderType: contracts.Desc,
	})
	return this
}
