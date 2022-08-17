package api

import (
	"MAS/shared"
	"MAS/wasm/engine"
	"MAS/wasm/states"
	"MAS/wasm/utils"
	"fmt"
	"strconv"
	"strings"
	"syscall/js"
)

//GetMapsAPI gets the list of maps from the server
func GetMapsAPI(this js.Value, inputs []js.Value) interface{} {

	handler := func(results []js.Value) error {
		mapList := js.Global().Get("document").Call(shared.FUNC_GET_E, shared.ID_MAP_LIST)

		var mapListHTML strings.Builder
		mapListHTML.Grow(64)

		r := []shared.TypeMap{}
		if err := convertMessage(results, &r); err != nil {
			return err
		}

		if len(r) == 0 {
			mapList.Set("innerHTML", `<li class='list-group-item d-flex justify-content-between align-items-center'>No saved maps!</li>`)
			return nil
		}

		for _, result := range r {
			fmt.Fprintf(&mapListHTML, `<li class='list-group-item d-flex justify-content-between align-items-center'>
			ID:%d -- NAME:%s
			<button class="badge badge-danger badge-pill" onclick="DeleteMapByID('/maps/%d');">Delete</button></li>`, result.ID, result.Name, result.ID)
		}

		mapList.Set("innerHTML", mapListHTML.String())

		return nil
	}

	makeAPICallGet(inputs, handler)
	return nil
}

func DeleteMapByID(this js.Value, inputs []js.Value) interface{} {
	url := inputs[0].String()
	makeAPICallPost(url, shared.API_DELETE, nil)
	return nil
}

func GetMapByIDAPI(this js.Value, inputs []js.Value) interface{} {

	v := js.Global().Get("document").Call(shared.FUNC_GET_E, shared.ID_MAP_VIS).Get("value").String()
	id, convErr := strconv.Atoi(v)

	if convErr != nil {
		utils.ShowMessage("ID is not a number:" + convErr.Error())
		return nil
	}

	url := fmt.Sprintf("/maps/%d", id)

	handler := func(results []js.Value) error {
		r := shared.TypeMap{}
		if err := convertMessage(results, &r); err != nil {
			return err
		}

		engine.LoadMap(r)

		return nil
	}

	makeAPICallGet([]js.Value{js.ValueOf(url)}, handler)

	return nil
}

func PostMapAPI(this js.Value, inputs []js.Value) interface{} {
	if states.MapIsEmpty() {
		utils.ShowMessage("Cannot save an empty map.")
		return nil
	}

	makeAPICallPost(inputs[0].String(), shared.API_POST, states.MapGetState())
	return nil
}
