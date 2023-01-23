package blog_test

import (
	"errors"
	"files/blog"
	"io/fs"
	"reflect"
	"strings"
	"testing"
	"testing/fstest"
)

type failingFS struct{}

func (f failingFS) Open(n string) (fs.File, error) {
	return nil, errors.New("dictionary not found")
}

func TestFileList(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		t.Skip()

		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte("hi")},
			"hello-world2.md": {Data: []byte("hola")},
		}
		got, err := blog.NewPostsFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		if lg, lf := len(got), len(fs); lg != lf {
			t.Errorf("got %d posts, wanted %d posts", lg, lf)
		}
	})

	assertPost := func(t testing.TB, got, want blog.Post) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %+v, want: %+v", got, want)
		}
	}
	t.Run("should parse title", func(t *testing.T) {
		post1 := `Title: post 1
Description: this is a very interesting post
Tags: domain driven design,c#, .NET
---
Hello World!
blah blah blah`
		post2 := `Title: post 2
Description: this is a very interesting post
Tags: domain driven design, golang
---
Hello World!
blah blah blah`
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(post1)},
			"hello-world2.md": {Data: []byte(post2)},
		}
		posts, err := blog.NewPostsFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		got := posts[0]
		want := blog.Post{
			Title:       "post 1",
			Description: "this is a very interesting post",
			Tags:        []string{"domain driven design", "c#", ".NET"},
			Body: `Hello World!
blah blah blah`,
		}
		assertPost(t, got, want)

	})

	t.Run("handle an error when reading files", func(t *testing.T) {

		_, err := blog.NewPostsFromFS(failingFS{})

		if err == nil {
			t.Error("expected an error")
		}

		if !strings.Contains(err.Error(), "new posts from fs") {
			t.Errorf("expected error to be wrapped")
		}
	})
}
