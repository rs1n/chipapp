package global

import "github.com/globalsign/mgo"

var global *Global

func GetGlobal() *Global {
	if global == nil {
		panic("global.GetGlobal: global is not initialized")
	}
	return global
}

func InitGlobalFor(hrp HtmlRenderParams, mgoDialInfo *mgo.DialInfo) {
	if global != nil {
		panic("global.InitGlobalFor: global is already initialized")
	}
	global = NewGlobal(hrp, mgoDialInfo)
}
