package gotv

import (
	"bytes"
	"text/template"
)

var (
	index = bytes.NewBuffer([]byte{})
)

func html() {
	buf := bytes.NewBuffer([]byte{})

	template.Must(template.ParseFiles(flags.js)).Execute(buf, struct {
		Host      string
		Ws        string
		Wid       int
		Hei       int
		EnableZip bool
	}{flags.host + ":" + flags.port, flags.ws, flags.width, flags.height, flags.enable_zip})

	template.Must(template.New("index.html").Parse(`<html><head><title>{{.Title}}</title></head><body><script>{{ .Js }}</script></body></html>`)).Execute(index, struct {
		Title string
		Js    string
	}{flags.title, buf.String()})
}
