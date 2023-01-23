package renderer_test

import (
	"bytes"
	"files/blog"
	"files/renderer"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRenderer(t *testing.T) {
	t.Run("Render a blog post", func(t *testing.T) {
		post1 := blog.Post{
			Title:       "post 1",
			Description: "this is a very interesting post",
			Tags:        []string{"domain driven design", "c#", ".NET"},
			Body: `# Hello World!
## some subchapter 1
- some list element 1
- some list element 2`,
		}

		var output bytes.Buffer

		pr, err := renderer.NewPostRenderer()
		assertNoError(t, err)
		err = pr.RenderPost(&output, post1)
		assertNoError(t, err)

		approvals.VerifyWithExtension(t, &output, ".html")
	})
	t.Run("Render an index", func(t *testing.T) {
		posts := []blog.Post{
			{Title: "Hello World"},
			{Title: "Hello World 2"},
		}

		var output bytes.Buffer
		pr, err := renderer.NewPostRenderer()
		assertNoError(t, err)
		err = pr.RenderIndex(&output, posts)
		assertNoError(t, err)

		approvals.VerifyWithExtension(t, &output, ".html")

	})

	t.Run("Sanitize title", func(t *testing.T) {
		want := "hello-world-2"
		got := renderer.SanitizeTitle("Hello World 2")

		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	})

}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

// func BenchmarkRenderer(b *testing.B) {
// 	post1 := blog.Post{
// 		Title:       "post 1",
// 		Description: "this is a very interesting post",
// 		Tags:        []string{"domain driven design", "c#", ".NET"},
// 		Body: `Hello World!
// 		blah blah blah`,
// 	}

// 	pr, _ := renderer.NewPostRenderer()
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		pr.Render(io.Discard, post1)
// 	}
// }
