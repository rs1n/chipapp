package global

var global *Global

func GetGlobal() *Global {
	if global == nil {
		panic("global.GetGlobal: global is not initialized")
	}
	return global
}

func InitGlobalFor(rhp HtmlRendererParams) {
	if global != nil {
		panic("global.InitGlobalFor: global is already initialized")
	}
	global = NewGlobal(rhp)
}
