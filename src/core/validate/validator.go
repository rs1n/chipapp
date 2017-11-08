package validate

import "github.com/rs1n/chipapp/src/core/global"

func Struct(s interface{}) error {
	g := global.GetGlobal()
	return g.Validate.Struct(s)
}
