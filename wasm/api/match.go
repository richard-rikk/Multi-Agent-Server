package api

import (
	"MAS/shared"
	"MAS/wasm/engine"
	"MAS/wasm/utils"
	"fmt"
	"strconv"
	"syscall/js"
)

func GetMatchByIDAPI(this js.Value, inputs []js.Value) interface{} {
	v := js.Global().Get("document").Call(shared.FUNC_GET_E, shared.ID_MATCH_ID).Get("value").String()
	id, convErr := strconv.Atoi(v)
	url := fmt.Sprintf("/match/%d", id)

	if convErr != nil {
		utils.ShowMessage("ID is not a number:" + convErr.Error())
		return nil
	}

	handler := func(results []js.Value) error {
		r := shared.TypeMatch{}
		if err := convertMessage(results, &r); err != nil {
			return err
		}

		engine.LoadMatch(r)
		return nil
	}

	makeAPICallGet([]js.Value{js.ValueOf(url)}, handler)

	return nil
}
