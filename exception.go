package querybuilder

import (
	"github.com/goal-web/contracts"
)

type ParamException struct {
	Err       error
	Arg       any
	Condition string
	previous  contracts.Exception
}

func (p ParamException) Error() string {
	return p.Err.Error()
}

func (p ParamException) GetPrevious() contracts.Exception {
	return p.previous
}
