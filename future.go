package querybuilder

import "github.com/goal-web/contracts"

func (builder *Builder) SetExecutor(executor contracts.SqlExecutor) contracts.QueryBuilder {
	return builder.QueryBuilder.SetExecutor(executor)
}

func (builder *Builder) Create(fields contracts.Fields) interface{} {
	return builder.QueryBuilder.Create(fields)
}

func (builder *Builder) Insert(values ...contracts.Fields) bool {
	return builder.QueryBuilder.Insert(values...)
}

func (builder *Builder) Delete() int64 {
	return builder.QueryBuilder.Delete()
}

func (builder *Builder) Update(fields contracts.Fields) int64 {
	return builder.QueryBuilder.Update(fields)
}

func (builder *Builder) Get() contracts.Collection {
	return builder.QueryBuilder.Get()
}

func (builder *Builder) SelectForUpdate() contracts.Collection {
	return builder.QueryBuilder.SelectForUpdate()
}

func (builder *Builder) Find(key interface{}) interface{} {
	return builder.QueryBuilder.Find(key)
}

func (builder *Builder) First() interface{} {
	return builder.QueryBuilder.First()
}

func (builder *Builder) InsertGetId(values ...contracts.Fields) int64 {
	return builder.QueryBuilder.InsertGetId(values...)
}

func (builder *Builder) InsertOrIgnore(values ...contracts.Fields) int64 {
	return builder.QueryBuilder.InsertOrIgnore(values...)
}

func (builder *Builder) InsertOrReplace(values ...contracts.Fields) int64 {
	return builder.QueryBuilder.InsertOrReplace(values...)
}

func (builder *Builder) FirstOrCreate(values ...contracts.Fields) interface{} {
	return builder.QueryBuilder.FirstOrCreate(values...)
}

func (builder *Builder) UpdateOrInsert(attributes contracts.Fields, values ...contracts.Fields) bool {
	return builder.QueryBuilder.UpdateOrInsert(attributes, values...)
}

func (builder *Builder) UpdateOrCreate(attributes contracts.Fields, values ...contracts.Fields) interface{} {
	return builder.QueryBuilder.UpdateOrCreate(attributes, values...)
}

func (builder *Builder) FirstOrFail() interface{} {
	return builder.QueryBuilder.FirstOrFail()
}

func (builder *Builder) Count(columns ...string) int64 {
	return builder.QueryBuilder.Count(columns...)
}

func (builder *Builder) Avg(column string, as ...string) int64 {
	return builder.QueryBuilder.Avg(column, as...)
}

func (builder *Builder) Sum(column string, as ...string) int64 {
	return builder.QueryBuilder.Sum(column, as...)
}

func (builder *Builder) Max(column string, as ...string) int64 {
	return builder.QueryBuilder.Max(column, as...)
}

func (builder *Builder) Min(column string, as ...string) int64 {
	return builder.QueryBuilder.Min(column, as...)
}

func (builder *Builder) SimplePaginate(perPage int64, current ...int64) contracts.Collection {
	return builder.WithPagination(perPage, current...).Get()
}

func (builder *Builder) FirstOr(provider contracts.InstanceProvider) interface{} {
	if result := builder.First(); result != nil {
		return result
	}
	return provider()
}

func (builder *Builder) FirstWhere(column string, args ...interface{}) interface{} {
	return builder.Where(column, args...).First()
}

func (builder *Builder) Paginate(perPage int64, current ...int64) (contracts.Collection, int64) {
	return builder.SimplePaginate(perPage, current...), builder.Count()
}
