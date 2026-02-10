package handlers

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/fdanctl/p5r-stats/src/models"
	"github.com/fdanctl/p5r-stats/src/render"
	"github.com/fdanctl/p5r-stats/src/services"
)

func RadarHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	udata, err := services.ReadUserData()
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	stats := services.ComputeStats(udata.Activities)

	// Example data (0–100 scale)
	data := []services.Metric{
		{Label: "Knowledge", Value: float64(stats[models.Knowledge])},
		{Label: "Guts", Value: float64(stats[models.Guts])},
		{Label: "Proficiency", Value: float64(stats[models.Proficiency])},
		{Label: "Kindness", Value: float64(stats[models.Kindness])},
		{Label: "Charm", Value: float64(stats[models.Charm])},
	}

	svg := services.BuildRadarSVG(400, 400, data, -30, 100)
	fmt.Fprint(w, svg)
}
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
			return
		}
		slices.Reverse(userData.Activities)
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
