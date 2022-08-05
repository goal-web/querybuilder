package querybuilder

import (
	"fmt"
)

func (this *Builder[T]) DeleteSql() (sql string, bindings []any) {
	sql = fmt.Sprintf("delete from %s", this.table)

	if !this.wheres.IsEmpty() {
		sql = fmt.Sprintf("%s where %s", sql, this.wheres.String())
	}
	bindings = this.GetBindings()
	return
}
