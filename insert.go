package querybuilder

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"strings"
)

func getInsertType(insertType2 ...contracts.InsertType) contracts.InsertType {
	if len(insertType2) > 0 {
		return insertType2[0]
	}
	return contracts.Insert
}

func (builder *Builder) CreateSql(value contracts.Fields, insertType2 ...contracts.InsertType) (sql string, bindings []interface{}) {
	if len(value) == 0 {
		return
	}
	keys := make([]string, 0)

	valuesString := fmt.Sprintf("(%s)", strings.Join(utils.MakeSymbolArray("?", len(value)), ","))
	for name, field := range value {
		bindings = append(bindings, field)
		keys = append(keys, name)
	}

	sql = fmt.Sprintf("%s into %s %s values %s", getInsertType(insertType2...), builder.table, fmt.Sprintf("(%s)", strings.Join(keys, ",")), valuesString)
	return
}

func (builder *Builder) InsertSql(values []contracts.Fields, insertType2 ...contracts.InsertType) (sql string, bindings []interface{}) {
	if len(values) == 0 {
		return
	}
	fields := utils.GetMapKeys(values[0])
	valuesString := make([]string, 0)

	for _, value := range values {
		valuesString = append(valuesString, fmt.Sprintf("(%s)", strings.Join(utils.MakeSymbolArray("?", len(value)), ",")))
		for _, field := range fields {
			bindings = append(bindings, value[field])
		}
	}

	fieldsString := fmt.Sprintf(" (%s)", strings.Join(fields, ","))

	sql = fmt.Sprintf("%s into %s%s values %s", getInsertType(insertType2...), builder.table, fieldsString, strings.Join(valuesString, ","))
	return
}

func (builder *Builder) InsertIgnoreSql(values []contracts.Fields) (sql string, bindings []interface{}) {
	return builder.InsertSql(values, contracts.InsertIgnore)
}

func (builder *Builder) InsertReplaceSql(values []contracts.Fields) (sql string, bindings []interface{}) {
	return builder.InsertSql(values, contracts.InsertReplace)
}
