package render

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/fdanctl/p5r-stats/src/utils"
)

var templates *template.Template

func Init() {
	templates = template.Must(
		template.New("").Funcs(template.FuncMap{
			"upper":      utils.Upper,
			"capitalize": utils.Capitalize,
			"titleCase":  utils.ToTitleCase,
		}).ParseGlob("src/templates/*.html"),
	)
}

func HTML(w http.ResponseWriter, name string, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := templates.ExecuteTemplate(w, name, data)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		http.Error(w, "Template error", http.StatusInternalServerError)
	}
}
