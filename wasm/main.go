package main

import (
	"MAS/wasm/api"
	"MAS/wasm/engine"
	"MAS/wasm/states"
	"MAS/wasm/utils"
	"syscall/js"
)

func registerCallbacks() {
	js.Global().Set("validateInput", js.FuncOf(utils.ValidateInput))
	js.Global().Set("getSettingsAPI", js.FuncOf(api.GetSettingsAPI))
	js.Global().Set("postSettingsAPI", js.FuncOf(api.PostSettingsAPI))
	js.Global().Set("RenderEmptyMap", js.FuncOf(engine.RenderEmptyMap))
	js.Global().Set("ChangeUI", js.FuncOf(utils.ChangeUI))
	js.Global().Set("getMapsAPI", js.FuncOf(api.GetMapsAPI))
	js.Global().Set("MapStateSetActive", js.FuncOf(states.MapSetActive))
	js.Global().Set("ModifyCell", js.FuncOf(engine.ModifyCell))
	js.Global().Set("PostMapAPI", js.FuncOf(api.PostMapAPI))
	js.Global().Set("GetMapByIDAPI", js.FuncOf(api.GetMapByIDAPI))
	js.Global().Set("DeleteMapByID", js.FuncOf(api.DeleteMapByID))
	js.Global().Set("GetMatchByIDAPI", js.FuncOf(api.GetMatchByIDAPI))
	js.Global().Set("Draw", js.FuncOf(utils.Draw))

}

func main() {

	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}
