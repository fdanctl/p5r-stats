package render

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/fdanctl/p5r-stats/config"
	"github.com/fdanctl/p5r-stats/src/utils"
)

type page uint8

const (
	PageHome page = iota
	Page404
	PageNewUser
	PageDesignSystem
	PageTest
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
		"dict":       utils.Dict,
	}

	globalPartialsSrc := fmt.Sprint(config.TmplsFolder, "partials/global/")
	globalPartialsDir, err := os.ReadDir(globalPartialsSrc)
	if err != nil {
		panic(err)
	}

	var globalPartials []string
	for _, v := range globalPartialsDir {
		globalPartials = append(
			globalPartials,
			fmt.Sprint(globalPartialsSrc, v.Name()),
		)
	}

	pages = map[page]mapValue{
		PageHome: {
			tmpl: template.Must(
				template.New("").Funcs(funcs).ParseFiles(
					append([]string{
						"src/templates/layouts/base.html",
						"src/templates/pages/home.html",
					}, globalPartials...)...,
				)),
			entry: "base.html",
		},
		PageNewUser: {
			tmpl: template.Must(
				template.New("").Funcs(funcs).ParseFiles(
					append([]string{
						"src/templates/layouts/base.html",
						"src/templates/pages/new.html",
					}, globalPartials...)...,
				)),
			entry: "base.html",
		},
		PageDesignSystem: {
			tmpl: template.Must(
				template.New("").Funcs(funcs).ParseFiles(
					"src/templates/layouts/base.html",
					"src/templates/pages/design-system.html",
				)),
			entry: "base.html",
		},
		PageTest: {
			tmpl: template.Must(
				template.New("").Funcs(funcs).ParseFiles(
					append([]string{
						"src/templates/layouts/base.html",
						"src/templates/pages/test.html",
					}, globalPartials...)...,
				)),
			entry: "base.html",
		},
		Page404: {
			tmpl: template.Must(
				template.ParseFiles(
					"src/templates/layouts/base.html",
					"src/templates/pages/404.html",
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
