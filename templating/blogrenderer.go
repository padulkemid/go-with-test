package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	"strings"
)

//go:embed "templates/*"
var postTemplates embed.FS

const (
	templateFiles = "templates/*.html"
	blogFile      = "blog.html"
)

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

type Renderer struct {
	templ *template.Template
}

func NewRenderer() (*Renderer, error) {
	templ, err := template.ParseFS(postTemplates, templateFiles)
	if err != nil {
		return nil, err
	}

	return &Renderer{templ: templ}, nil
}

func (r *Renderer) Render(w io.Writer, p Post) error {
	if err := r.templ.ExecuteTemplate(w, blogFile, p); err != nil {
		return err
	}

	return nil
}

func (r *Renderer) RenderIndex(w io.Writer, p []Post) error {
	indexTemplate := `<ol>{{range .}}<li><a href="/post/{{sanitiseTitle .Title}}">{{.Title}}</a></li>{{end}}</ol>`
	sanitiseFunc := template.FuncMap{
		"sanitiseTitle": func(title string) string {
			return strings.ToLower(strings.Replace(title, " ", "-", -1))
		},
	}

	templ, err := template.New("index").Funcs(sanitiseFunc).Parse(indexTemplate)
	if err != nil {
		return err
	}

	if err := templ.Execute(w, p); err != nil {
		return err
	}

	return nil
}
