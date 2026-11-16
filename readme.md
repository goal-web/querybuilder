# Goal/QueryBuilder
Goal 的数据库查询构造器为创建和运行数据库查询提供了一个方便的接口。它可以用于支持大部分数据库操作，并与 Goal 支持的所有数据库系统完美运行。并且大量参考了 `Laravel` 的查询构造器设计，你几乎可以在这个库找到所有与 `Laravel` 对应的方法。

Goal 的查询构造器实现了类似 PDO 参数绑定的形式，来保护您的应用程序免受 SQL 注入攻击。因此不必清理因参数绑定而传入的字符串。查询构造器会返回你想要的 SQL 语句以及绑定参数。

## 运行数据库查询
### 根据条件从表中检索出数据
你可以使用 `New()` 方法来开始查询。该方法为给定的表返回一个查询构造器实例，允许你在查询上链式调用更多的约束，最后使用 get 方法获取结果：

### 查询语句
```go
package tests

import (
	"fmt"
	"github.com/goal-web/contracts"
	builder "github.com/goal-web/querybuilder"
	"github.com/stretchr/testify/assert"
	"github.com/xwb1989/sqlparser"
	"testing"
)
func TestSimpleQueryBuilder(t *testing.T) {
    query := builder.New[contracts.Fields]("users")
    query.Where("name", "qbhy").
    Where("age", ">", 18).
    Where("gender", "!=", 0, contracts.Or).
    OrWhere("amount", ">=", 100).
    WhereIsNull("avatar")
    fmt.Println(query.ToSql())
    // select * from `users` where name = ? and age > ? and avatar is null or gender != ? or amount >= ?
    fmt.Println(query.GetBindings())
    // [qbhy 18 0 100]
    
    _, err := sqlparser.Parse(query.ToSql())
    assert.Nil(t, err, err)
}
```

### 插入语句
你可以通过 `InsertSql` 或者 `CreateSql` 很方便的生成插入语句。
```go
package tests

import (
	"fmt"
	"github.com/goal-web/contracts"
	builder "github.com/goal-web/querybuilder"
	"github.com/stretchr/testify/assert"
	"github.com/xwb1989/sqlparser"
	"testing"
)

func TestInsertSql(t *testing.T) {
	sql, bindings := builder.New[contracts.Fields]("users").InsertSql([]contracts.Fields{
		{"name": "qbhy", "age": 18, "money": 100000000000, "gender": nil},
		{"name": "goal", "age": 18, "money": 10, "gender": nil},
	})
	fmt.Println(sql)
	fmt.Println(bindings)
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}
```
 
### 更新语句
你可以通过 `UpdateSql` 很方便的生成更新语句。
```go
package tests

import (
	"fmt"
	"github.com/goal-web/contracts"
	builder "github.com/goal-web/querybuilder"
	"github.com/stretchr/testify/assert"
	"github.com/xwb1989/sqlparser"
	"testing"
)

func TestUpdateSql(t *testing.T) {
	type Settings struct {
		Name string `json:"name"`
	}
	sql, bindings := builder.New[contracts.Fields]("users").Where("id", ">", 1).UpdateSql(contracts.Fields{
		"name": "qbhy", "age": 18, "money": 100000000000,
		"settings": Settings{Name: "json_name"},
	})
	fmt.Println(sql)
	fmt.Println(bindings)
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)

	rawUpdateSql, rawUpdateBindings := builder.New[contracts.Fields]("users").
		Where("id", ">", 1).
		UpdateSql(contracts.Fields{
			"name":  "qbhy",
			"age":   builder.Expression("`age` + 10"),
			"money": 100000000000,
		})
	fmt.Println(rawUpdateSql)
	fmt.Println(rawUpdateBindings)
	_, rawUpdateErr := sqlparser.Parse(rawUpdateSql)
	assert.Nil(t, rawUpdateErr, rawUpdateErr)
}
```

### 删除语句
你可以通过 `DeleteSql` 很方便的生成删除语句。
```go
package tests

import (
	"fmt"
	"github.com/goal-web/contracts"
	builder "github.com/goal-web/querybuilder"
	"github.com/stretchr/testify/assert"
	"github.com/xwb1989/sqlparser"
	"testing"
)

func TestDeleteSql(t *testing.T) {
	sql, bindings := builder.New[contracts.Fields]("users").Where("id", ">", 1).DeleteSql()
	fmt.Println(sql)
	fmt.Println(bindings)
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}
```

## 更多高级用法
支持 where嵌套、子查询、连表、连子查询等更多高级用法

正如开头所说，你可以在这里找到几乎所有与 `Laravel` 对应的查询构造器方法，也可以在 [测试文件](tests/querybuilder_test.go) 中找到更多用法

[goal/query-builder](https://github.com/goal-web/querybuilder)  
qbhy0715@qq.com
