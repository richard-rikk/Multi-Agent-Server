package utils

import (
	"MAS/shared"
	"syscall/js"
)

func Draw(this js.Value, inputs []js.Value) interface{} {
	plot := js.Global().Get(shared.V_PLOT)
	canvas := js.Global().Get("document").Call(shared.FUNC_GET_E, shared.ID_GRAPH_RESRC)
	data := map[string]interface{}{"y": []interface{}{1, 2, 4, 8, 16}}
	setting := map[string]interface{}{"margin": map[string]interface{}{"t": 0}}

	plot.Call("newPlot", canvas,
		[]interface{}{data},
		setting,
	)

	return nil
}
