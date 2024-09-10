package web

import (
	"bytes"
	"testing"
)

func TestRender(t *testing.T) {
	var (
		aPost = Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)
	t.Run("convert single post to HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		result := buf.String()
		expected := `<h1>hello world</h1>
<p>This is a description</p>
Tags: <ul><li>go</li><li>tdd</li></ul>`

		if result != expected {
			t.Errorf("result '%s' expected '%s'", result, expected)
		}
	})
}
