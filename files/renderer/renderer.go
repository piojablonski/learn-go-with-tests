package renderer

import (
	"embed"
	"files/blog"
	"fmt"
	"html/template"
	"io"
	"strings"

	"github.com/gomarkdown/markdown"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type PostRenderer struct {
	template *template.Template
}

type PostViewModel struct {
	blog.Post
	BodyAsHtml     template.HTML
	SanitizedTitle string
}

func ParseBody(body string) template.HTML {

	parsed := markdown.ToHTML([]byte(body), nil, nil)
	return template.HTML(string(parsed))

}

func ConvertPostToViewModel(post blog.Post) (PostViewModel, error) {
	model := PostViewModel{
		Post:           post,
		BodyAsHtml:     ParseBody(post.Body),
		SanitizedTitle: SanitizeTitle(post.Title),
	}
	return model, nil
}

func SanitizeTitle(in string) string {
	return strings.ToLower(strings.ReplaceAll(in, " ", "-"))
}

func NewPostRenderer() (*PostRenderer, error) {

	templ1, err := template.New("").ParseFS(postTemplates, "templates/*.gohtml")

	if err != nil {
		return nil, err
	}
	return &PostRenderer{template: templ1}, nil
}

func (pr *PostRenderer) RenderPost(w io.Writer, post blog.Post) error {

	viewModel, err := ConvertPostToViewModel(post)
	if err != nil {
		return fmt.Errorf("rendering post failed %w", err)
	}
	err = pr.template.ExecuteTemplate(w, "post.gohtml", viewModel)
	if err != nil {
		return err
	}

	return nil

}
func (pr *PostRenderer) RenderIndex(wr io.Writer, posts []blog.Post) error {
	tmp, err := template.New("index.gohtml").ParseFS(postTemplates, "templates/*")

	if err != nil {
		return err
	}

	var postViewModels []PostViewModel
	for _, p := range posts {
		vm, err := ConvertPostToViewModel(p)
		if err != nil {
			return err
		}
		postViewModels = append(postViewModels, vm)

	}

	err = tmp.ExecuteTemplate(wr, "index.gohtml", postViewModels)
	if err != nil {
		return err
	}

	return nil

}
