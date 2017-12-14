package global

import "upper.io/db.v3/postgresql"

var global *Global

func GetGlobal() *Global {
	if global == nil {
		panic("global.GetGlobal: global is not initialized")
	}
	return global
}

func InitGlobalFor(
	hrp HtmlRenderParams, connectionURL *postgresql.ConnectionURL,
) {
	if global != nil {
		panic("global.InitGlobalFor: global is already initialized")
	}
	global = NewGlobal(hrp, connectionURL)
}
