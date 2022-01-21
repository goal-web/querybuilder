package querybuilder

import "github.com/goal-web/contracts"

func (this *Builder) SetTX(tx interface{}) contracts.QueryBuilder {
	//TODO implement me
	panic("implement me")
}

func (this *Builder) Paginate(perPage int64, current ...int64) (interface{}, int64) {
	//TODO implement me
	panic("implement me")
}

func (this *Builder) Create(fields contracts.Fields) interface{} {
	//TODO implement me
	panic("implement me")
}

func (this *Builder) Insert(values ...contracts.Fields) bool {
	//TODO implement me
	panic("implement me")
}

func (this *Builder) Delete() int64 {
	//TODO implement me
	panic("implement me")
}

func (this *Builder) Update(fields contracts.Fields) int64 {
	//TODO implement me
	panic("implement me")
}

func (this *Builder) Get() interface{} {
	//TODO implement me
	panic("implement me")
}

func (this *Builder) Find(key interface{}) interface{} {
	//TODO implement me
	panic("implement me")
}

func (this *Builder) First() interface{} {
	//TODO implement me
	panic("implement me")
}

func (this *Builder) InsertGetId(values ...contracts.Fields) int64 {
	//TODO implement me
	panic("implement me")
}

func (this *Builder) InsertOrIgnore(values ...contracts.Fields) int64 {
	//TODO implement me
	panic("implement me")
}

func (this *Builder) InsertOrReplace(values ...contracts.Fields) int64 {
	//TODO implement me
	panic("implement me")
}

func (this *Builder) FirstOrCreate(values ...contracts.Fields) interface{} {
	//TODO implement me
	panic("implement me")
}

func (this *Builder) UpdateOrInsert(attributes contracts.Fields, values ...contracts.Fields) bool {
	//TODO implement me
	panic("implement me")
}

func (this *Builder) UpdateOrCreate(attributes contracts.Fields, values ...contracts.Fields) interface{} {
	//TODO implement me
	panic("implement me")
}

func (this *Builder) FirstOr(provider contracts.InstanceProvider) interface{} {
	//TODO implement me
	panic("implement me")
}

func (this *Builder) FirstOrFail(provider contracts.InstanceProvider) interface{} {
	//TODO implement me
	panic("implement me")
}

func (this *Builder) FirstWhere(column string, args ...interface{}) interface{} {
	//TODO implement me
	panic("implement me")
}

func (this *Builder) Count(columns ...string) int64 {
	//TODO implement me
	panic("implement me")
}

func (this *Builder) Avg(column string, as ...string) int64 {
	//TODO implement me
	panic("implement me")
}

func (this *Builder) Sum(column string, as ...string) int64 {
	//TODO implement me
	panic("implement me")
}

func (this *Builder) Max(column string, as ...string) int64 {
	//TODO implement me
	panic("implement me")
}

func (this *Builder) Min(column string, as ...string) int64 {
	//TODO implement me
	panic("implement me")
}
