package web

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io"
)

//# impl #2
//const (
//	postTemplate = `<h1>{{.Title}}</h1><p>{{.Description}}</p>Tags: <ul>{{range .Tags}}<li>{{.}}</li>{{end}}</ul>`
//)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func Render(w io.Writer, p Post) error {
	//impl #3
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	//impl #2
	/*	templ, err := template.New("post").Parse(postTemplate)
	 */
	if err != nil {
		return err
	}
	if err = templ.Execute(w, p); err != nil {
		return err
	}

	//impl #1
	//_, err := fmt.Fprintf(w, "<h1>%s</h1>\n<p>%s</p>\n%s", p.Title, p.Description, buildTags(p))
	//return err
	return nil
}

func buildTags(p Post) string {
	var buff = bytes.Buffer{}
	_, _ = buff.Write([]byte("Tags: <ul>"))
	for _, tag := range p.Tags {
		fmt.Fprintf(&buff, "<li>%s</li>", tag)
	}
	buff.Write([]byte("</ul>"))
	return buff.String()
}
