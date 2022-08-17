package utils

import (
	"MAS/shared"
	"strconv"
	"syscall/js"
)

func ShowMessage(msg string) {
	window := js.Global().Get("window")
	window.Call("alert", msg)
}

// Add a list of CSS classes to an element.
func addClasses(element js.Value, classes []string) {
	for _, c := range classes {
		element.Get("classList").Call("add", c)
	}
}

// Clears any CSS class in the list from element.
func clearClasses(element js.Value, classes []string) {
	for _, c := range classes {
		element.Get("classList").Call("remove", c)
	}
}

// Validate number inputs between range.
// inputs =
// 0 -> string	id of the element to be validated
// 1 -> int 	the value has to be greater than this number
// 2 -> int 	the value has to be less than this number
func ValidateInput(this js.Value, inputs []js.Value) interface{} {
	document := js.Global().Get("document")
	element := document.Call(shared.FUNC_GET_E, inputs[0].String())
	low := inputs[1].Int()
	high := inputs[2].Int()
	valid := "is-valid"
	invalid := "is-invalid"

	v, convErr := strconv.Atoi(element.Get("value").String())
	clearClasses(element, []string{valid, invalid})

	if convErr != nil || v < low || v > high {
		addClasses(element, []string{invalid})
		return false
	}

	addClasses(element, []string{valid})
	return true

}

//Used for changing the UI properties like color and tile sizes.
func ChangeUI(this js.Value, inputs []js.Value) interface{} {
	id := inputs[0].String()
	indID := inputs[1].String()
	cssProperty := inputs[2].String()
	property := inputs[3].String()
	value := js.Global().Get("document").Call("getElementById", id).Get("value").String()

	if value == "" {
		return nil
	}

	root := js.Global().Get("document").Get("documentElement")
	root.Get("style").Call("setProperty", cssProperty, value)

	if property != "" {
		indicator := js.Global().Get("document").Call("getElementById", indID)
		indicator.Get("style").Set(property, value)
	}

	return nil
}
