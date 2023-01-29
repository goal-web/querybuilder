package querybuilder

import (
	"fmt"
)

func (builder *Builder) DeleteSql() (sql string, bindings []interface{}) {
	sql = fmt.Sprintf("delete from %s", builder.table)

	if !builder.wheres.IsEmpty() {
		sql = fmt.Sprintf("%s where %s", sql, builder.wheres.String())
	}
	bindings = builder.GetBindings()
	return
}
