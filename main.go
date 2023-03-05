package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"github.com/creack/pty"
	"github.com/webview/webview"
)

type PtyRespEvent struct {
	Result string `json:"result"`
}

func main() {
	c := exec.Command("bash")
	p, _ := pty.Start(c)

	w := webview.New(true)
	defer w.Destroy()

	w.SetTitle("Ratatouille")
	w.SetSize(720, 432, webview.HintFixed)

	b, _ := ioutil.ReadFile("frontend/dist/index.html")
	w.SetHtml(string(b))

	w.Bind("__WRITE_PTY", func(command string) {
		p.WriteString(command + "\n")
	})

	go func() {
		time.Sleep(500 * time.Millisecond)
		for {
			buf := make([]byte, 1024)
			n, err := p.Read(buf)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}

			result := string(buf)
			event := PtyRespEvent{Result: result[:n]}
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
