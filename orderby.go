package querybuilder

import (
	"fmt"
	"github.com/goal-web/contracts"
	"strings"
)

const RandomOrder contracts.OrderType = "RANDOM()"
const RandOrder contracts.OrderType = "RAND()"

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
		if orderBy.field == "" {
			columns = append(columns, fmt.Sprintf("%s", orderBy.fieldOrderType))
		} else {
			columns = append(columns, fmt.Sprintf("%s %s", orderBy.field, orderBy.fieldOrderType))
		}
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

func (this *Builder) InRandomOrder() contracts.QueryBuilder {
	this.orderBy = append(this.orderBy, OrderBy{
		field:          "",
		fieldOrderType: RandomOrder,
	})
	return this
}
