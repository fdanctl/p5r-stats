package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/fdanctl/p5r-stats/src/models"
	"github.com/fdanctl/p5r-stats/src/render"
	"github.com/fdanctl/p5r-stats/src/services"
)

type iname struct {
	Name string
}

func UserFormHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/partials/user/edit/"):]

	switch r.Method {

	case http.MethodGet:
		render.HTML(w, render.FragmentUserEdit, iname{name})

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func UserInfoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/partials/user/edit-cancel/"):]

	switch r.Method {

	case http.MethodGet:
		render.HTML(w, render.FragmentUserEditCancel, iname{name})

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}


func UserDataHandler(w http.ResponseWriter, r *http.Request) {
switch r.Method {
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Bad Form", http.StatusBadRequest)
			return
		}

		userData, err := services.CreateUserData(r.Form.Get("name"))
		if err != nil {
			fmt.Println(err)
			if errors.Is(err, models.ErrAlreadyExists) {
				http.Error(w, "User data found.\nTry reloading the page.", http.StatusBadRequest)
				return
			}
		}

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
		
		render.HTML(w, render.FragmentHome, data)
	}
}
