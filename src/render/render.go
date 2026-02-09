package render

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/fdanctl/p5r-stats/config"
	"github.com/fdanctl/p5r-stats/src/utils"
)

type view uint8

const (
	PageHome view = iota
	Page404
	PageDesignSystem
	PageTest

	FragmentHome
	FragmentUserEdit
	FragmentUsernameDiv
	FragmentModal
	FragmentStatSelect
	FragmentActivity
	FragmentStatsGraph

	FragmentToast
)

type mapValue struct {
	tmpl  *template.Template
	entry string
}

var pages map[view]mapValue

func Init() {
	funcs := template.FuncMap{
		"upper":        strings.ToUpper,
		"capitalize":   utils.Capitalize,
		"titleCase":    utils.ToTitleCase,
		"timeToString": utils.TimeToString,
		"dict":         utils.Dict,
		"randInt":      rand.Intn,
		"parseStat":    utils.StatToString,
	}

	globalPartialsSrc := fmt.Sprint(config.TmplsFolder, "partials/global/")
	globalPartialsDir, err := os.ReadDir(globalPartialsSrc)
	if err != nil {
		panic(err)
	}

	var globalPartials []string
	for _, v := range globalPartialsDir {
		fmt.Println(globalPartialsSrc, v.Name())
		globalPartials = append(
			globalPartials,
			fmt.Sprint(globalPartialsSrc, v.Name()),
		)
	}

	pages = map[view]mapValue{
		PageHome: {
			tmpl: template.Must(
				template.New("").Funcs(funcs).ParseFiles(
					append([]string{
						"src/templates/partials/features/activity-card.html",
						"src/templates/partials/features/profile-header.html",
						"src/templates/partials/features/stats-graph.html",
						"src/templates/layouts/base.html",
						"src/templates/pages/home.html",
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

		FragmentHome: {
			tmpl: template.Must(
				template.New("").Funcs(funcs).ParseFiles(
					append([]string{
						"src/templates/partials/features/activity-card.html",
						"src/templates/partials/features/profile-header.html",
						"src/templates/partials/features/stats-graph.html",
						"src/templates/pages/home.html",
					}, globalPartials...)...,
				)),
			entry: "content",
		},
		FragmentUsernameDiv: {
			tmpl: template.Must(
				template.New("").Funcs(funcs).ParseFiles(
					append([]string{
						"src/templates/partials/features/profile-header.html",
					}, globalPartials...)...,
				)),
			entry: "username",
		},
		FragmentUserEdit: {
			tmpl: template.Must(
				template.New("").Funcs(funcs).ParseFiles(
					append([]string{
						"src/templates/partials/features/profile-header.html",
					}, globalPartials...)...,
				)),
			entry: "editing",
		},
		FragmentModal: {
			tmpl: template.Must(
				template.New("").Funcs(funcs).ParseFiles(
					globalPartials...,
				)),
			entry: "modal.html",
		},
		FragmentStatSelect: {
			tmpl: template.Must(
				template.New("").Funcs(funcs).ParseFiles(
					globalPartials...,
				)),
			entry: "select-stat",
		},
		FragmentActivity: {
			tmpl: template.Must(
				template.New("").Funcs(funcs).ParseFiles(
					append([]string{
						"src/templates/partials/features/activity-card.html",
					}, globalPartials...)...,
				)),
			entry: "activity-card",
		},
		FragmentStatsGraph: {
			tmpl: template.Must(
				template.New("").Funcs(funcs).ParseFiles(
					append([]string{
						"src/templates/partials/features/stats-graph.html",
					}, globalPartials...)...,
				)),
			entry: "stats-graph",
		},
		FragmentToast: {
			tmpl: template.Must(
				template.New("").Funcs(funcs).ParseFiles(
					globalPartials...,
				)),
			entry: "toast",
		},
	}
}

type OOB struct {
	ID   string
	Swap string
	View view
	Data any
}

func HTML(w http.ResponseWriter, view view, data any, oob []OOB) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl, entry := pages[view].tmpl, pages[view].entry
	err := tmpl.ExecuteTemplate(w, entry, data)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		http.Error(w, "Template error", http.StatusInternalServerError)
	}

	for _, v := range oob {
		RenderOOB(w, v.ID, v.Swap, v.View, v.Data)
	}
}

func RenderOOB(w http.ResponseWriter, id, swap string, view view, data any) {
	fmt.Fprintf(w, `<div id="%s" hx-swap-oob="%s">`, id, swap)
	tmpl, entry := pages[view].tmpl, pages[view].entry
	err := tmpl.ExecuteTemplate(w, entry, data)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		http.Error(w, "Template error", http.StatusInternalServerError)
	}
	fmt.Fprint(w, `</div>`)
}
