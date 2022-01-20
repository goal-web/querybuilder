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

func (this *Builder) Insert(values ...contracts.Fields) interface{} {
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
