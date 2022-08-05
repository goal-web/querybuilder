package querybuilder

import "github.com/goal-web/contracts"

func (this *Builder[T]) SetExecutor(executor contracts.SqlExecutor[T]) contracts.QueryBuilder[T] {
	return this.QueryBuilder.SetExecutor(executor)
}

func (this *Builder[T]) Create(fields contracts.Fields) T {
	return this.QueryBuilder.Create(fields)
}

func (this *Builder[T]) Insert(values ...contracts.Fields) bool {
	return this.QueryBuilder.Insert(values...)
}

func (this *Builder[T]) Delete() int64 {
	return this.QueryBuilder.Delete()
}

func (this *Builder[T]) Update(fields contracts.Fields) int64 {
	return this.QueryBuilder.Update(fields)
}

func (this *Builder[T]) Get() contracts.Collection[T] {
	return this.QueryBuilder.Get()
}

func (this *Builder[T]) SelectForUpdate() contracts.Collection[T] {
	return this.QueryBuilder.SelectForUpdate()
}

func (this *Builder[T]) Find(key any) (T, bool) {
	return this.QueryBuilder.Find(key)
}

func (this *Builder[T]) First() (T, bool) {
	return this.QueryBuilder.First()
}

func (this *Builder[T]) InsertGetId(values ...contracts.Fields) int64 {
	return this.QueryBuilder.InsertGetId(values...)
}

func (this *Builder[T]) InsertOrIgnore(values ...contracts.Fields) int64 {
	return this.QueryBuilder.InsertOrIgnore(values...)
}

func (this *Builder[T]) InsertOrReplace(values ...contracts.Fields) int64 {
	return this.QueryBuilder.InsertOrReplace(values...)
}

func (this *Builder[T]) FirstOrCreate(values ...contracts.Fields) T {
	return this.QueryBuilder.FirstOrCreate(values...)
}

func (this *Builder[T]) UpdateOrInsert(attributes contracts.Fields, values ...contracts.Fields) bool {
	return this.QueryBuilder.UpdateOrInsert(attributes, values...)
}

func (this *Builder[T]) UpdateOrCreate(attributes contracts.Fields, values ...contracts.Fields) T {
	return this.QueryBuilder.UpdateOrCreate(attributes, values...)
}

func (this *Builder[T]) FirstOrFail() T {
	return this.QueryBuilder.FirstOrFail()
}

func (this *Builder[T]) Count(columns ...string) int64 {
	return this.QueryBuilder.Count(columns...)
}

func (this *Builder[T]) Avg(column string, as ...string) int64 {
	return this.QueryBuilder.Avg(column, as...)
}

func (this *Builder[T]) Sum(column string, as ...string) int64 {
	return this.QueryBuilder.Sum(column, as...)
}

func (this *Builder[T]) Max(column string, as ...string) int64 {
	return this.QueryBuilder.Max(column, as...)
}

func (this *Builder[T]) Min(column string, as ...string) int64 {
	return this.QueryBuilder.Min(column, as...)
}

func (this *Builder[T]) SimplePaginate(perPage int64, current ...int64) contracts.Collection[T] {
	return this.WithPagination(perPage, current...).Get()
}

func (this *Builder[T]) FirstOr(provider func() T) T {
	if result, exists := this.First(); exists {
		return result
	}
	return provider()
}

func (this *Builder[T]) FirstWhere(column string, args ...any) (T, bool) {
	return this.Where(column, args...).First()
}

func (this *Builder[T]) Paginate(perPage int64, current ...int64) (contracts.Collection[T], int64) {
	return this.SimplePaginate(perPage, current...), this.Count()
}
