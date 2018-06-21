// +build js,wasm

package webapi

import "syscall/js"

// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget
type EventTarget interface {
	AddEventListener(typ string, listener func(args []js.Value))
	RemoveEventListener()
	DispatchEvent()
}

type implEventTarget struct {
	v js.Value
}

func (r *implEventTarget) AddEventListener(typ string, listener func(args []js.Value)) {
	r.v.Call("addEventListener", typ, js.NewCallback(listener))
}

func (r *implEventTarget) RemoveEventListener() {
	panic("RemoveEventListener")
}

func (r *implEventTarget) DispatchEvent() {
	panic("DispatchEvent")
}
