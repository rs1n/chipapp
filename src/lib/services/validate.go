package services

import "github.com/sknv/chipapp/src/core/global"

type Validate struct{}

func (_ *Validate) ValidateStruct(s interface{}) error {
	g := global.GetGlobal()
	return g.Validate.Struct(s)
}
