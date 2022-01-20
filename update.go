package querybuilder

import (
	"fmt"
	"github.com/goal-web/contracts"
	"strings"
)

func (this *Builder) UpdateSql(value contracts.Fields) (sql string, bindings []interface{}) {
	if len(value) == 0 {
		return
	}
	valuesString := make([]string, 0)
	for name, field := range value {
		valuesString = append(valuesString, fmt.Sprintf("%s = ?", name))
		bindings = append(bindings, field)
	}

	sql = fmt.Sprintf("update %s set %s", this.table, strings.Join(valuesString, ","))

	if !this.wheres.IsEmpty() {
		sql = fmt.Sprintf("%s where %s", sql, this.wheres.String())
	}

	bindings = append(bindings, this.bindings[whereBinding]...)

	return
}
