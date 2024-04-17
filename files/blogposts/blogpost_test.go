package blogpost_test

import (
	"errors"
	blogpost "hello/files"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("opened twice, go to fail")
}

func assertPost(t testing.TB, got, want blogpost.Post) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestNewBlogPosts(t *testing.T) {
	t.Run("success opening file", func(t *testing.T) {
		t.Skip()
		fs := fstest.MapFS{
			"test.md": {
				Data: []byte("Title: Post 1"),
			},
			"test_2.md": {
				Data: []byte("Title: Post 2"),
			},
		}

		posts, err := blogpost.NewPostFromFS(fs)
		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fs) {
			t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
		}
	})

	t.Run("failed opening file", func(t *testing.T) {
		fs := StubFailingFS{}
		_, err := blogpost.NewPostFromFS(fs)

		if err == nil {
			t.Fatal("should return an error!")
		}
	})

	t.Run("success checking metadata", func(t *testing.T) {
		md := `Title: Post 1
Description: Description 1
Tags: tag1, tag2
----
this is the
body`

		fs := fstest.MapFS{
			"test_meta.md": {
				Data: []byte(md),
			},
		}

		posts, _ := blogpost.NewPostFromFS(fs)

		got := posts[0]
		want := blogpost.Post{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"tag1", "tag2"},
			Body: `this is the
body`,
		}

		assertPost(t, got, want)
	})
}
