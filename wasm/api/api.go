package api

import (
	"MAS/wasm/utils"
	"encoding/json"
	"fmt"
	"syscall/js"
)

//convertMessage converts the response of the server and serializes
//the data into the container.
//
//It returns an error if json.Unmarshal returns an error.
//
//Use this function in Get function calls.
func convertMessage(results []js.Value, container interface{}) error {
	Json := js.Global().Get("JSON")
	return json.Unmarshal([]byte(Json.Call("stringify", results[0]).String()), container)
}

// makeAPICallGet use this function to handle API Get calls inside the Go runtime.
// Makes a Javascript Fetch call using the URL in inputs[0]. Shows the errors to the users.
// The handler function needs to handle the result given by the server and return any error so it can be shown to the user.
// The function will run in a seperate Goroutines so it wont block the main event loop.
func makeAPICallGet(inputs []js.Value, handler func([]js.Value) error) {
	fetch := js.Global().Get("fetch")

	url := inputs[0].String()
	awaitable := fetch.Invoke(url)
	ch := make(chan []js.Value)
	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ch <- args
		return nil
	})

	awaitable.Call("then", cb)
	go func() {
		results := <-ch
		response := results[0]

		if !(response.Get("ok").Bool()) {
			utils.ShowMessage("Server error code: " + response.Get("status").String())
			return
		}

		awaitable = response.Call("json")
		go awaitable.Call("then", cb)
		results = <-ch

		if err := handler(results); err != nil {
			utils.ShowMessage(err.Error())
			return
		}

	}()

}

// makeAPICallPost makes a post request to the API server with the given body.
// It can handle UPDATE and DELETE methods as well using the method parameter.
func makeAPICallPost(url, method string, body interface{}) {
	fetch := js.Global().Get("fetch")
	b, err := json.Marshal(body)

	if err != nil {
		utils.ShowMessage(err.Error())
		return
	}

	responseBody := map[string]interface{}{
		// Adding method type
		"method": method,

		// Adding body or contents to send
		"body": string(b),

		// Adding headers to the request
		"headers": map[string]interface{}{"Content-type": "application/json; charset=UTF-8"},
	}

	awaitable := fetch.Invoke(url, responseBody)
	ch := make(chan []js.Value)
	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ch <- args
		return nil
	})

	awaitable.Call("then", cb)
	go func() {
		results := <-ch
		response := results[0]

		if !(response.Get("ok").Bool()) {
			//utils.ShowMessage("Server error code: " + response.Get("status").String())
			utils.ShowMessage(fmt.Sprintf("Server error code: %d, Error message: %s", response.Get("status").Int(), response.Get("statusText").String()))
			return
		}

		awaitable = response.Call("json")
		go awaitable.Call("then", cb)
		results = <-ch

		js.Global().Get("console").Call("log", results[0])
	}()

}
