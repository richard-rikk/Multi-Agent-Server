package api

import (
	"MAS/shared"
	"MAS/wasm/utils"
	"fmt"
	"strconv"
	"syscall/js"
)

//GetSettingsAPI gets the settings from the server and shows them on the UI.
func GetSettingsAPI(this js.Value, inputs []js.Value) interface{} {
	handler := func(results []js.Value) error {
		document := js.Global().Get("document")

		r := shared.TypeSettings{}
		if err := convertMessage(results, &r); err != nil {
			return err
		}

		for key, value := range r {
			document.Call(shared.FUNC_GET_E, key).Set("value", value)
		}

		return nil
	}

	makeAPICallGet(inputs, handler)

	return nil
}

//PostSettingsAPI sends an API Post request to the server and validates the values of the Game Settings form.
//The order of the values determined by their appearance on the Game Settings form.
func PostSettingsAPI(this js.Value, inputs []js.Value) interface{} {
	document := js.Global().Get("document")
	formElements := document.Call(shared.FUNC_GET_E, shared.ID_SET_GAME_FORM).Get("elements")
	n := document.Call(shared.FUNC_GET_E, shared.ID_SET_GAME_FORM).Get("elements").Get("length").Int()
	res := shared.TypeSettings{}

	for i := 0; i < n; i++ {
		item := formElements.Call("item", i)
		if item.Get("nodeName").String() == "INPUT" || item.Get("nodeName").String() == "SELECT" {
			v, convErr := strconv.ParseUint(item.Get("value").String(), 10, 0)

			if convErr != nil {
				utils.ShowMessage(fmt.Sprintf("Check %d. element of the form and try again. Error:%s", i+1, convErr.Error()))
				return false
			}
			res[item.Get("id").String()] = uint(v)
		}
	}

	makeAPICallPost(inputs[0].String(), shared.API_POST, res)
	return true
}
