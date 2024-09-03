package web

import "syscall/js"

var Window JSWindow

type JSWindow struct{}

var funcMaps = make(map[string]func(js.Value, []js.Value) any)

func (w *JSWindow) AddEventListener(inputEvent string, inputFunc func(js.Value, []js.Value) any) {
	js.Global().Get("window").Call("addEventListener", inputEvent, js.FuncOf(inputFunc))
	funcMaps[inputEvent] = inputFunc
}

func (w *JSWindow) RemoveEventListener(inputEvent string) {
	js.Global().Get("window").Call("removeEventListener", inputEvent, js.FuncOf(funcMaps[inputEvent]))
	delete(funcMaps, inputEvent)
}
