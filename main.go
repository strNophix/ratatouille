package main

import (
	"io/ioutil"
	"time"

	"github.com/webview/webview"
)

func main() {
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("Basic Example")
	w.SetSize(480, 320, webview.HintNone)

	b, _ := ioutil.ReadFile("frontend/dist/index.html")
	w.SetHtml(string(b))

	go func() {
		for {
			time.Sleep(2 * time.Second)
			w.Eval(`window.__WRITE_TERMINAL('PONG!\r\n');`)
		}
	}()

	w.Run()
}
