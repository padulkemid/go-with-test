package blogrenderer_test

import (
	"bytes"
	blogrenderer "hello/templating"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	post := blogrenderer.Post{
		Title:       "hello this is a title!",
		Body:        "this is a body",
		Description: "this is a description",
		Tags:        []string{"go", "tdd"},
	}

	renderer, err := blogrenderer.NewRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := renderer.Render(&buf, post); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})

	t.Run("renders index", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blogrenderer.Post{
			{Title: "title 1"},
			{Title: "title 2"},
		}

		if err := renderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<ol><li><a href="/post/title-1">title 1</a></li><li><a href="/post/title-2">title 2</a></li></ol>`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func BenchmarkRender(b *testing.B) {
	post := blogrenderer.Post{
		Title:       "hello this is a title!",
		Body:        "this is a body",
		Description: "this is a description",
		Tags:        []string{"go", "tdd"},
	}

	b.ResetTimer()

	renderer, err := blogrenderer.NewRenderer()
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		if err := renderer.Render(io.Discard, post); err != nil {
			b.Fatal("should not throw error")
		}
	}
}
