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
		Host string
		Ws   string
	}{flags.host, flags.ws})

	template.Must(template.New("index.html").Parse(`<html><head><title>{{.Title}}</title></head><body><script>{{ .Js }}</script></body></html>`)).Execute(index, struct {
		Title string
		Js    string
	}{flags.title, buf.String()})
}
