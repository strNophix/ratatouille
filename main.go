package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"time"

	"github.com/creack/pty"
	"github.com/webview/webview"
)

type PtyOutEvent struct {
	Result string `json:"result"`
}

func main() {
	c := exec.Command("bash")
	p, _ := pty.Start(c)

	w := webview.New(true)
	defer w.Destroy()

	w.SetTitle("Ratatouille")
	w.SetSize(480, 320, webview.HintNone)

	b, _ := ioutil.ReadFile("frontend/dist/index.html")
	w.SetHtml(string(b))

	w.Bind("__WRITE_PTY", func(command string) {
		p.WriteString(command + "\n")
	})

	go func() {
		for {
			buf := make([]byte, 1024)
			n, _ := p.Read(buf)
			if n == 0 {
				continue
			}

			result := string(buf)
			event := PtyOutEvent{Result: result[:n]}
			payload, err := json.Marshal(event)
			if err != nil {
				fmt.Println("Whoopsie", err)
				continue
			}

			fn := fmt.Sprintf(`window.__WRITE_TERMINAL(%s);`, string(payload))
			w.Eval(fn)
			time.Sleep(200 * time.Millisecond)
		}
	}()

	w.Run()
}
