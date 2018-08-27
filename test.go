// +build js,wasm

package main

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"syscall/js"
)

var send js.Value

func main() {
	href := js.Global().Get("location").Get("href")
	u, err := url.Parse(href.String())
	if err != nil {
		log.Fatal(err)
	}
	u.Path = "/logo.png"

	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	enc := base64.StdEncoding.EncodeToString(b)

	canvas := js.Global().Get("document").Call("getElementById", "canvas")
	ctx := canvas.Call("getContext", "2d")
	image := js.Global().Call("eval", "new Image()")
	image.Call("addEventListener", "load", js.NewCallback(func(args []js.Value) {
		ctx.Call("drawImage", image, 0, 0)
	}))
	image.Set("src", "data:image/png;base64,"+enc)

	canvas.Call("addEventListener", "click", js.NewCallback(func(args []js.Value) {
		js.Global().Get("window").Call("alert", "Don't click me!")
	}))

	onLoad()
	select {}
}

func onLoad() {
	// 要素に値を指定可能かつイベント登録可能
	input := js.Global().Get("document").Call("getElementById", "input")
	input.Set("value", "abc")
	input.Call("addEventListener", "change", js.NewCallback(onChange))

	// アラートを表示可能
	button := js.Global().Get("document").Call("getElementById", "button")
	button.Set("disabled", false)
	button.Call("addEventListener", "click", js.NewCallback(func(args []js.Value) {
		js.Global().Get("alert").Invoke("WebAssembly!")
	}))

	// Ajax 的な動作も可能
	send = js.Global().Get("document").Call("getElementById", "send")
	send.Set("disabled", false)
	send.Call("addEventListener", "click", js.NewCallback(func(args []js.Value) {
		send.Set("innerHTML", "Sending...")
		send.Set("disabled", true)
		go onClick()
	}))
}

func onChange(args []js.Value) {
	js.Global().Get("text").Set("innerHTML", args[0].Get("target").Get("value"))
}

func onClick() {
	u, err := url.Parse("https://httpbin.org")
	if err != nil {
		log.Fatal(err)
	}
	u.Path = "/get"

	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	js.Global().Get("text").Set("innerHTML", string(b))
	send.Set("innerHTML", "Send")
	send.Set("disabled", false)
}
