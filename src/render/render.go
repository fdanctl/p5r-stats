package render

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/fdanctl/p5r-stats/src/utils"
)

type page uint8

const (
	PageHome page = iota
	PageNewUser
	PageDesignSystem
)

type mapValue struct {
	tmpl  *template.Template
	entry string
}

var pages map[page]mapValue

func Init() {
	funcs := template.FuncMap{
		"upper":      utils.Upper,
		"capitalize": utils.Capitalize,
		"titleCase":  utils.ToTitleCase,
	}

	pages = map[page]mapValue{
		PageHome: {
			tmpl: template.Must(
				template.New("").Funcs(funcs).ParseFiles(
					"src/templates/base.html", "src/templates/home.html",
				)),
			entry: "base.html",
		},
		PageNewUser: {
			tmpl: template.Must(
				template.New("").Funcs(funcs).ParseFiles(
					"src/templates/base.html", "src/templates/new.html",
				)),
			entry: "base.html",
		},
		PageDesignSystem: {
			tmpl: template.Must(
				template.New("").Funcs(funcs).ParseFiles(
					"src/templates/base.html", "src/templates/design-system.html",
				)),
			entry: "base.html",
		},
	}
}

func HTML(w http.ResponseWriter, page page, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl, entry := pages[page].tmpl, pages[page].entry
	err := tmpl.ExecuteTemplate(w, entry, data)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		http.Error(w, "Template error", http.StatusInternalServerError)
	}
}
