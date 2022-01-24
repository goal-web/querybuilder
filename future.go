package querybuilder

import "github.com/goal-web/contracts"

func (this *Builder) SetExecutor(executor contracts.SqlExecutor) contracts.QueryBuilder {
	return this.QueryBuilder.SetExecutor(executor)
}

func (this *Builder) Create(fields contracts.Fields) interface{} {
	return this.QueryBuilder.Create(fields)
}

func (this *Builder) Insert(values ...contracts.Fields) bool {
	return this.QueryBuilder.Insert(values...)
}

func (this *Builder) Delete() int64 {
	return this.QueryBuilder.Delete()
}

func (this *Builder) Update(fields contracts.Fields) int64 {
	return this.QueryBuilder.Update(fields)
}

func (this *Builder) Get() contracts.DBCollection {
	return this.QueryBuilder.Get()
}

func (this *Builder) Find(key interface{}) interface{} {
	return this.QueryBuilder.Find(key)
}

func (this *Builder) First() interface{} {
	return this.QueryBuilder.First()
}

func (this *Builder) InsertGetId(values ...contracts.Fields) int64 {
	return this.QueryBuilder.InsertGetId(values...)
}

func (this *Builder) InsertOrIgnore(values ...contracts.Fields) int64 {
	return this.QueryBuilder.InsertOrIgnore(values...)
}

func (this *Builder) InsertOrReplace(values ...contracts.Fields) int64 {
	return this.QueryBuilder.InsertOrReplace(values...)
}

func (this *Builder) FirstOrCreate(values ...contracts.Fields) interface{} {
	return this.QueryBuilder.FirstOrCreate(values...)
}

func (this *Builder) UpdateOrInsert(attributes contracts.Fields, values ...contracts.Fields) bool {
	return this.QueryBuilder.UpdateOrInsert(attributes, values...)
}

func (this *Builder) UpdateOrCreate(attributes contracts.Fields, values ...contracts.Fields) interface{} {
	return this.QueryBuilder.UpdateOrCreate(attributes, values...)
}

func (this *Builder) FirstOrFail() interface{} {
	return this.QueryBuilder.FirstOrFail()
}

func (this *Builder) Count(columns ...string) int64 {
	return this.QueryBuilder.Count(columns...)
}

func (this *Builder) Avg(column string, as ...string) int64 {
	return this.QueryBuilder.Avg(column, as...)
}

func (this *Builder) Sum(column string, as ...string) int64 {
	return this.QueryBuilder.Sum(column, as...)
}

func (this *Builder) Max(column string, as ...string) int64 {
	return this.QueryBuilder.Max(column, as...)
}

func (this *Builder) Min(column string, as ...string) int64 {
	return this.QueryBuilder.Min(column, as...)
}

func (this *Builder) SimplePaginate(perPage int64, current ...int64) interface{} {
	return this.WithPagination(perPage, current...).Get()
}

func (this *Builder) FirstOr(provider contracts.InstanceProvider) interface{} {
	if result := this.First(); result != nil {
		return result
	}
	return provider()
}

func (this *Builder) FirstWhere(column string, args ...interface{}) interface{} {
	return this.Where(column, args...).First()
}

func (this *Builder) Paginate(perPage int64, current ...int64) (interface{}, int64) {
	return this.SimplePaginate(perPage, current...), this.Count()
}
