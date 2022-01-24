# Goal/QueryBuilder
Goal 的数据库查询构造器为创建和运行数据库查询提供了一个方便的接口。它可以用于支持大部分数据库操作，并与 Goal 支持的所有数据库系统完美运行。并且大量参考了 `Laravel` 的查询构造器设计，你几乎可以在这个库找到所有与 `Laravel` 对应的方法。

Goal 的查询构造器实现了类似 PDO 参数绑定的形式，来保护您的应用程序免受 SQL 注入攻击。因此不必清理因参数绑定而传入的字符串。查询构造器会返回你想要的 SQL 语句以及绑定参数。

## 运行数据库查询
### 根据条件从表中检索出数据
你可以使用 `NewQuery` 方法来开始查询。该方法为给定的表返回一个查询构造器实例，允许你在查询上链式调用更多的约束，最后使用 get 方法获取结果：
```golang
package querybuilder
import (
	"fmt"
)

func TestSimpleQueryBuilder() {
    query := NewQuery("users").
		Where("name", "qbhy").
        Where("age", ">", 18).
        Where("gender", "!=", 0).
        OrWhere("amount", ">=", 100).
        WhereIsNull("avatar")
	
    fmt.Println(query.ToSql())
    fmt.Println(query.GetBindings())
    // select * from users where name = ? and age > ? and gender != ? and avatar is null or amount >= ?
    // [qbhy 18 0 100]
}
```
> 你也可以通过 `SelectSql` 方法一次性获取你想要的参数。
> 例如：sql, bindings := NewQuery("users").Where("gender", 1).SelectSql()

### 插入语句
你可以通过 `InsertSql` 或者 `CreateSql` 很方便的生成插入语句。
```golang
package querybuilder

import (
	"fmt"
	"github.com/goal-web/contracts"
)

// TestInsertSql 批量插入数据
func TestInsertSql() {
	sql, bindings := NewQuery("users").InsertSql([]contracts.Fields{
		{"name": "qbhy", "age": 18, "money": 100000000000},
		{"name": "goal", "age": 18, "money": 10},
	})
	fmt.Println(sql)
	fmt.Println(bindings)
	// insert into users (name,age,money) values (?,?,?),(?,?,?)
	// [qbhy 18 100000000000 goal 18 10]
}
// TestCreateSql 插入单个数据
func TestCreateSql() {
	sql, bindings := NewQuery("users").CreateSql(contracts.Fields{
		"name": "qbhy", "age": 18, "money": 100000000000,
	})
	fmt.Println(sql)
	fmt.Println(bindings) 
	// insert into users (name,age,money) values (?,?,?) 
	//[qbhy 18 100000000000]
}
```
 
### 更新语句
你可以通过 `UpdateSql` 很方便的生成更新语句。
```golang
package querybuilder

import (
	"fmt"
	"github.com/goal-web/contracts"
)

func TestUpdateSql() {
	sql, bindings := NewQuery("users").Where("id", ">", 1).UpdateSql(contracts.Fields{
		"name": "qbhy", "age": 18, "money": 100000000000,
	})
	fmt.Println(sql)
	fmt.Println(bindings)
    // update users set money = ?,name = ?,age = ? where id > ?
    // [qbhy 18 100000000000 1]
}
```

### 删除语句
你可以通过 `DeleteSql` 很方便的生成删除语句。
```golang
package querybuilder

import (
	"fmt"
)

func TestDeleteSql() {
	sql, bindings := NewQuery("users").Where("id", ">", 1).DeleteSql()
	fmt.Println(sql)
	fmt.Println(bindings)
    // delete from users where id > ?
    // [1]
}
```

## 更多高级用法
支持 where嵌套、子查询、连表、连子查询等更多高级用法
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
	query := builder.NewQuery("users")
	query.Where("name", "qbhy").
		Where("age", ">", 18).
		Where("gender", "!=", 0, contracts.Or).
		OrWhere("amount", ">=", 100).
		WhereIsNull("avatar")
	fmt.Println(query.ToSql())
	fmt.Println(query.GetBindings())

	_, err := sqlparser.Parse(query.ToSql())
	assert.Nil(t, err, err)
}

func TestJoinQueryBuilder(t *testing.T) {
	query := builder.NewQuery("users").
		Join("accounts", "accounts.user_id", "=", "users.id").
		JoinSub(func() contracts.QueryBuilder {
			return builder.NewQuery("users").
				Where("level", ">", 5)
		}, "vip_users", "vip_users.id", "=", "users.id").
		//WhereIn("gender", "1,2").
		WhereIn("gender", []int{1, 2})
	fmt.Println(query.ToSql())
	fmt.Println(query.GetBindings())
	_, err := sqlparser.Parse(query.ToSql())
	assert.Nil(t, err, err)
}

func TestFromSubQueryBuilder(t *testing.T) {
	query := builder.FromSub(func() contracts.QueryBuilder {
		return builder.NewQuery("users").
			Where("level", ">", 5)
	}, "vip_users").
		//WhereIn("gender", "1,2").
		WhereIn("gender", []int{1, 2})
	fmt.Println(query.ToSql())
	fmt.Println(query.GetBindings())
	_, err := sqlparser.Parse(query.ToSql())
	assert.Nil(t, err, err)
}

func TestDistinctQueryBuilder(t *testing.T) {
	query := builder.NewQuery("users").
		Distinct().
		Join("accounts", "accounts.user_id", "=", "users.id").
		Where("gender", "!=", 0, contracts.Or)
	fmt.Println(query.ToSql())
	fmt.Println(query.GetBindings())
	_, err := sqlparser.Parse(query.ToSql())
	assert.Nil(t, err, err)
}

func TestUpdateSql(t *testing.T) {
	sql, bindings := builder.NewQuery("users").Where("id", ">", 1).UpdateSql(contracts.Fields{
		"name": "qbhy", "age": 18, "money": 100000000000,
	})
	fmt.Println(sql)
	fmt.Println(bindings)
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}
func TestSelectSub(t *testing.T) {
	sql, bindings := builder.NewQuery("users").Where("id", ">", 1).
		SelectSub(func() contracts.QueryBuilder {
			return builder.NewQuery("accounts").Where("accounts.id", "users.id").WithCount()
		}, "accounts_count").
		Join("accounts", "accounts.user_id", "=", "users.id").
		SelectSql()
	fmt.Println(sql)
	fmt.Println(bindings)
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}
func TestWhereNotExists(t *testing.T) {
	sql, bindings := builder.NewQuery("users").
		Where("id", ">", 1).
		WhereNotExists(func() contracts.QueryBuilder {
			return builder.NewQuery("users").Select("id").Where("age", ">", 18)
		}).
		SelectSql()
	fmt.Println(sql)
	fmt.Println(bindings)
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}
func TestCount(t *testing.T) {
	sql, bindings := builder.NewQuery("users").Where("id", ">", 1).WithCount("id").SelectSql()
	fmt.Println(sql)
	fmt.Println(bindings)
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}
func TestDeleteSql(t *testing.T) {
	sql, bindings := builder.NewQuery("users").Where("id", ">", 1).DeleteSql()
	fmt.Println(sql)
	fmt.Println(bindings)
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}
func TestInsertSql(t *testing.T) {
	sql, bindings := builder.NewQuery("users").InsertSql([]contracts.Fields{
		{"name": "qbhy", "age": 18, "money": 100000000000},
		{"name": "goal", "age": 18, "money": 10},
	})
	fmt.Println(sql)
	fmt.Println(bindings)
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}
func TestInsertIgnoreSql(t *testing.T) {
	sql, bindings := builder.NewQuery("users").InsertIgnoreSql([]contracts.Fields{
		{"name": "qbhy", "age": 18, "money": 100000000000},
		{"name": "goal", "age": 18, "money": 10},
	})
	fmt.Println(sql)
	fmt.Println(bindings)
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}
func TestInsertReplaceSql(t *testing.T) {
	sql, bindings := builder.NewQuery("users").InsertReplaceSql([]contracts.Fields{
		{"name": "qbhy", "age": 18, "money": 100000000000},
		{"name": "goal", "age": 18, "money": 10},
	})
	fmt.Println(sql)
	fmt.Println(bindings)
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}

func TestCreateSql(t *testing.T) {
	sql, bindings := builder.NewQuery("users").CreateSql(contracts.Fields{
		"name": "qbhy", "age": 18, "money": 100000000000,
	})
	fmt.Println(sql)
	fmt.Println(bindings)
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}

func TestBetweenQueryBuilder(t *testing.T) {
	query := builder.NewQuery("users").
		Join("accounts", "accounts.user_id", "=", "users.id").
		WhereFunc(func(b contracts.QueryBuilder) {
			// 高瘦
			b.WhereBetween("height", []int{180, 200}).
				WhereBetween("weight", []int{50, 60}).
				WhereIn("id", []int{1, 2, 3, 4, 5})
		}).OrWhereFunc(func(b contracts.QueryBuilder) {
		// 矮胖
		b.WhereBetween("height", []int{140, 160}).
			WhereBetween("weight", []int{70, 140}).
			WhereNotBetween("id", []int{1, 5})
	})
	fmt.Println(query.ToSql())
	fmt.Println(query.GetBindings())
	_, err := sqlparser.Parse(query.ToSql())
	assert.Nil(t, err, err)
}

func TestUnionQueryBuilder(t *testing.T) {
	query := builder.NewQuery("users").
		Join("accounts", "accounts.user_id", "=", "users.id").
		Where("gender", "!=", 0, contracts.Or).
		UnionByProvider(
			func() contracts.QueryBuilder {
				return builder.NewQuery("peoples").Where("id", 5)
			},
		).
		Union(
			builder.NewQuery("accounts"),
		).
		UnionAll(
			builder.NewQuery("members"),
		).
		UnionAll(
			builder.NewQuery("students"),
		)
	fmt.Println(query.ToSql())
	fmt.Println(query.GetBindings())
	_, err := sqlparser.Parse(query.ToSql())
	assert.Nil(t, err, err)
}

func TestComplexQueryBuilder(t *testing.T) {

	query := builder.NewQuery("users")
	query.
		FromSub(func() contracts.QueryBuilder {
			return builder.NewQuery("users").Where("amount", ">", 1000)
		}, "rich_users").
		Join("accounts", "users.id", "=", "accounts.user_id").
		WhereFunc(func(b contracts.QueryBuilder) {
			b.Where("name", "goal").
				Where("age", "<", "18").
				WhereIn("id", []int{1, 2})
		}).
		OrWhereFunc(func(b contracts.QueryBuilder) {
			b.Where("name", "qbhy").
				Where("age", ">", 18).
				WhereNotIn("id", []int{1, 2})
		}).
		OrWhereNotIn("id", []int{6, 7}).
		OrWhereNotNull("id").
		OrderByDesc("age").
		OrderBy("id").
		GroupBy("country")

	fmt.Println(query.ToSql())
	fmt.Println(query.GetBindings())
	_, err := sqlparser.Parse(query.ToSql())
	assert.Nil(t, err, err)
}

func TestGroupByQueryBuilder(t *testing.T) {

	query := builder.
		FromSub(func() contracts.QueryBuilder {
			return builder.NewQuery("users").Where("amount", ">", 1000)
		}, "rich_users").
		GroupBy("country").
		Having("count(rich_users.id)", "<", 1000).   // 人口少
		OrHaving("sum(rich_users.amount)", "<", 100) // 或者穷

	fmt.Println(query.ToSql())
	fmt.Println(query.GetBindings())
	_, err := sqlparser.Parse(query.ToSql())
	assert.Nil(t, err, err)
}
```
正如开头所说，你可以在这里找到几乎所有与 `Laravel` 对应的查询构造器方法，也可以在 [测试文件](https://github.com/goal-web/querybuilder/blob/master/database_test.go) 中找到更多用法

[goal/query-builder](https://github.com/goal-web/querybuilder)  
qbhy0715@qq.com
