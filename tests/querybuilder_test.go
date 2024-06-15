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
	fmt.Println(query.GetBindings())

	_, err := sqlparser.Parse(query.ToSql())
	assert.Nil(t, err, err)
}

func TestJoinQueryBuilder(t *testing.T) {
	query := builder.New[contracts.Fields]("users").
		Join("accounts", "accounts.user_id", "=", "users.id").
		JoinSub(func() contracts.Query[contracts.Fields] {
			return builder.New[contracts.Fields]("users").
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
	query := builder.FromSub[contracts.Fields](func() contracts.Query[contracts.Fields] {
		return builder.New[contracts.Fields]("users").
			Where("level", ">", 5)
	}, "vip_users").
		//WhereIn("gender", "1,2").
		WhereIn("gender", []int{1, 2})
	sql := query.ToSql()
	fmt.Println(sql)
	fmt.Println(query.GetBindings())
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}

func TestDistinctQueryBuilder(t *testing.T) {
	query := builder.New[contracts.Fields]("users").
		Distinct().
		Join("accounts", "accounts.user_id", "=", "users.id").
		Where("gender", "!=", 0, contracts.Or)
	fmt.Println(query.ToSql())
	fmt.Println(query.GetBindings())
	_, err := sqlparser.Parse(query.ToSql())
	assert.Nil(t, err, err)
}

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

func TestSelectSub(t *testing.T) {
	sql, bindings := builder.New[contracts.Fields]("users").Where("id", ">", 1).
		SelectSub(func() contracts.Query[contracts.Fields] {
			return builder.New[contracts.Fields]("accounts").Where("accounts.id", "users.id").WithCount()
		}, "accounts_count").
		Join("accounts", "accounts.user_id", "=", "users.id").
		SelectSql()
	fmt.Println(sql)
	fmt.Println(bindings)
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}

func TestWhereByExpression(t *testing.T) {
	sql, bindings := builder.New[contracts.Fields]("users").
		Where("id", ">", 1).
		WhereIn("user_id", builder.Expression("(select user_id from follows where follower_id=1)")).
		SelectSql()
	fmt.Println(sql)
	fmt.Println(bindings)
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}

func TestWhereByQuery(t *testing.T) {
	sql, bindings := builder.New[contracts.Fields]("users").
		Where("id", ">", 1).
		WhereIn("user_id", builder.New[contracts.Fields]("follows").Where("follower_id", 2)).
		SelectSql()
	fmt.Println(sql)
	fmt.Println(bindings)
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}

func TestSelectForUpdate(t *testing.T) {
	sql, bindings := builder.New[contracts.Fields]("users").
		Where("id", ">", 1).
		SelectForUpdateSql()
	fmt.Println(sql)
	fmt.Println(bindings)
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}

func TestWhereNotExists(t *testing.T) {
	sql, bindings := builder.New[contracts.Fields]("users").
		Where("id", ">", 1).
		WhereNotExists(func() contracts.Query[contracts.Fields] {
			return builder.New[contracts.Fields]("users").Select("id").Where("age", ">", 18)
		}).
		SelectSql()
	fmt.Println(sql)
	fmt.Println(bindings)
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}
func TestCount(t *testing.T) {
	sql, bindings := builder.New[contracts.Fields]("users").Where("id", ">", 1).WithCount("id").SelectSql()
	fmt.Println(sql)
	fmt.Println(bindings)
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}
func TestDeleteSql(t *testing.T) {
	sql, bindings := builder.New[contracts.Fields]("users").Where("id", ">", 1).DeleteSql()
	fmt.Println(sql)
	fmt.Println(bindings)
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}
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
func TestInsertIgnoreSql(t *testing.T) {
	sql, bindings := builder.New[contracts.Fields]("users").InsertIgnoreSql([]contracts.Fields{
		{"name": "qbhy", "age": 18, "money": 100000000000},
		{"name": "goal", "age": 18, "money": 10},
	})
	fmt.Println(sql)
	fmt.Println(bindings)
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}
func TestInsertReplaceSql(t *testing.T) {
	sql, bindings := builder.New[contracts.Fields]("users").InsertReplaceSql([]contracts.Fields{
		{"name": "qbhy", "age": 18, "money": 100000000000},
		{"name": "goal", "age": 18, "money": 10},
	})
	fmt.Println(sql)
	fmt.Println(bindings)
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}

func TestCreateSql(t *testing.T) {
	sql, bindings := builder.New[contracts.Fields]("users").CreateSql(contracts.Fields{
		"name": "qbhy", "age": 18, "money": 100000000000,
	})
	fmt.Println(sql)
	fmt.Println(bindings)
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}

func TestBetweenQueryBuilder(t *testing.T) {
	query := builder.New[contracts.Fields]("users").
		Join("accounts", "accounts.user_id", "=", "users.id").
		WhereFunc(func(b contracts.Query[contracts.Fields]) {
			// 高瘦
			b.WhereBetween("height", []int{180, 200}).
				WhereBetween("weight", []int{50, 60}).
				WhereIn("id", []int{1, 2, 3, 4, 5})
		}).OrWhereFunc(func(b contracts.Query[contracts.Fields]) {
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
	query := builder.New[contracts.Fields]("users").
		Join("accounts", "accounts.user_id", "=", "users.id").
		Where("gender", "!=", 0, contracts.Or).
		UnionByProvider(func() contracts.Query[contracts.Fields] {
			return builder.New[contracts.Fields]("peoples").Where("id", 5)
		}).
		Union(
			builder.New[contracts.Fields]("accounts"),
		).
		UnionAll(
			builder.New[contracts.Fields]("members"),
		).
		UnionAll(
			builder.New[contracts.Fields]("students"),
		)
	fmt.Println(query.ToSql())
	fmt.Println(query.GetBindings())
	_, err := sqlparser.Parse(query.ToSql())
	assert.Nil(t, err, err)
}

func TestComplexQueryBuilder(t *testing.T) {

	query := builder.New[contracts.Fields]("users")
	query.
		FromSub(func() contracts.Query[contracts.Fields] {
			return builder.New[contracts.Fields]("users").Where("amount", ">", 1000)
		}, "rich_users").
		Join("accounts", "users.id", "=", "accounts.user_id").
		WhereFunc(func(b contracts.Query[contracts.Fields]) {
			b.Where("name", "goal").
				Where("age", "<", "18").
				WhereIn("id", []int{1, 2})
		}).
		OrWhereFunc(func(b contracts.Query[contracts.Fields]) {
			b.Where("name", "qbhy").
				Where("age", ">", 18).
				WhereNotIn("id", []int{1, 2})
		}).
		OrWhereNotIn("id", []int{6, 7}).
		OrWhereNotNull("id").
		OrderByDesc("age").
		OrderBy("id").
		GroupBy("country")

	sql, bindings := query.SelectSql()
	fmt.Println(sql)
	fmt.Println(bindings)
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}

func TestGroupByQueryBuilder(t *testing.T) {
	query := builder.
		FromSub[contracts.Fields](func() contracts.Query[contracts.Fields] {
		return builder.New[contracts.Fields]("users").Where("amount", ">", 1000)
	}, "rich_users").
		GroupBy("country").
		Having("count(rich_users.id)", "<", 1000).   // 人口少
		OrHaving("sum(rich_users.amount)", "<", 100) // 或者穷

	fmt.Println(query.ToSql())
	fmt.Println(query.GetBindings())
	_, err := sqlparser.Parse(query.ToSql())
	assert.Nil(t, err, err)
}

func TestInRandomOrder(t *testing.T) {
	query := builder.
		New[contracts.Fields]("users").
		GroupBy("country").
		Having("count(rich_users.id)", "<", 1000).    // 人口少
		OrHaving("sum(rich_users.amount)", "<", 100). // 或者穷
		InRandomOrder()

	fmt.Println(query.ToSql())
	fmt.Println(query.GetBindings())
	_, err := sqlparser.Parse(query.ToSql())
	assert.Nil(t, err, err)
}

func TestWhereIn(t *testing.T) {
	query := builder.
		New[contracts.Fields]("users").
		WhereNotIn("id", []any{1, 2, 3, 4}).
		InRandomOrder()

	sql := query.ToSql()
	fmt.Println(sql)
	fmt.Println(query.GetBindings())
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}

func TestWhen(t *testing.T) {
	query := builder.
		New[contracts.Fields]("users").
		When(true, func(q contracts.Query[contracts.Fields]) contracts.Query[contracts.Fields] {
			return q.Where("status", 1)
		}).
		WhereNotIn("id", []any{1, 2, 3, 4}).
		InRandomOrder()

	sql := query.ToSql()
	fmt.Println(sql)
	fmt.Println(query.GetBindings())
	_, err := sqlparser.Parse(sql)
	assert.Nil(t, err, err)
}
