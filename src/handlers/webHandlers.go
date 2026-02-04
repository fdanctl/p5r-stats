package handlers

import (
	"fmt"
	"net/http"

	"github.com/fdanctl/p5r-stats/src/models"
	"github.com/fdanctl/p5r-stats/src/render"
	"github.com/fdanctl/p5r-stats/src/services"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		render.HTML(w, render.PageTest, nil)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		if r.URL.Path != "/" {
			render.HTML(w, render.Page404, nil)
			// http.NotFound(w, r)
			return
		}

		userData, err := services.ReadUserData()

		if err != nil {
			render.HTML(w, render.PageHome, nil)
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
			render.HTML(w, render.PageHome, data)
		}

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func DesignHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		render.HTML(w, render.PageDesignSystem, nil)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

type iname struct {
	Name string
}

func UserFormHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/user/edit/"):]
	fmt.Printf("name: %v\n", name)

	switch r.Method {

	case http.MethodGet:
		render.HTML(w, render.FragmentUserEdit, iname{name})

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func UserInfoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/user/edit-cancel/"):]
	fmt.Printf("name: %v\n", name)

	switch r.Method {

	case http.MethodGet:
		render.HTML(w, render.FragmentUserEditCancel, iname{name})

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
