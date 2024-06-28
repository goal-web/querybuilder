package querybuilder

import (
	"encoding/json"
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"reflect"
	"strings"
)

func getInsertType(insertType2 ...contracts.InsertType) contracts.InsertType {
	if len(insertType2) > 0 {
		return insertType2[0]
	}
	return contracts.Insert
}

func wrapperValue(value any) any {
	if value == nil {
		return nil
	}
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Map, reflect.Struct, reflect.Array, reflect.Slice:
		jsonBytes, _ := json.Marshal(value)
		return string(jsonBytes)
	default:
		return value
	}
}

func (builder *Builder[T]) CreateSql(value contracts.Fields, insertType2 ...contracts.InsertType) (sql string, bindings []any) {
	if len(value) == 0 {
		return
	}
	keys := make([]string, 0)

	valuesString := fmt.Sprintf("(%s)", strings.Join(utils.MakeSymbolArray("?", len(value)), ","))
	for name, field := range value {
		bindings = append(bindings, wrapperValue(field))
		keys = append(keys, name)
	}

	sql = fmt.Sprintf("%s into `%s` %s values %s", getInsertType(insertType2...), builder.table, fmt.Sprintf("(%s)", strings.Join(keys, ",")), valuesString)
	return
}

func (builder *Builder[T]) InsertSql(values []contracts.Fields, insertType2 ...contracts.InsertType) (sql string, bindings []any) {
	if len(values) == 0 {
		return
	}
	fields := utils.GetMapKeys(values[0])
	valuesString := make([]string, 0)

	for _, value := range values {
		valuesString = append(valuesString, fmt.Sprintf("(%s)", strings.Join(utils.MakeSymbolArray("?", len(value)), ",")))
		for _, field := range fields {
			bindings = append(bindings, wrapperValue(value[field]))
		}
	}

	fieldsString := fmt.Sprintf(" (%s)", strings.Join(fields, ","))

	sql = fmt.Sprintf("%s into `%s`%s values %s", getInsertType(insertType2...), builder.table, fieldsString, strings.Join(valuesString, ","))
	return
}

func (builder *Builder[T]) InsertIgnoreSql(values []contracts.Fields) (sql string, bindings []any) {
	return builder.InsertSql(values, contracts.InsertIgnore)
}

func (builder *Builder[T]) InsertReplaceSql(values []contracts.Fields) (sql string, bindings []any) {
	return builder.InsertSql(values, contracts.InsertReplace)
}
