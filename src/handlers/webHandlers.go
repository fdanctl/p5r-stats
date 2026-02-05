package handlers

import (
	"net/http"

	"github.com/fdanctl/p5r-stats/src/models"
	"github.com/fdanctl/p5r-stats/src/render"
	"github.com/fdanctl/p5r-stats/src/services"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		render.HTML(w, render.PageTest, nil, nil)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		if r.URL.Path != "/" {
			render.HTML(w, render.Page404, nil, nil)
			// http.NotFound(w, r)
			return
		}

		userData, err := services.ReadUserData()

		if err != nil {
			render.HTML(w, render.PageHome, nil, nil)
		} else {
			stats := services.ComputeStats(userData.Activities)

			data := models.HomePageData{
				UserData: *userData,
				Stats: models.Stats{
					Knowledge:   stats[models.Knowledge],
					Guts:        stats[models.Guts],
					Proficiency: stats[models.Proficiency],
					Kindness:    stats[models.Kindness],
					Charm:       stats[models.Charm],
				},
			}
			render.HTML(w, render.PageHome, data, nil)
		}

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func DesignHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		render.HTML(w, render.PageDesignSystem, nil, nil)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
