package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/fdanctl/p5r-stats/src/models"
	"github.com/fdanctl/p5r-stats/src/render"
	"github.com/fdanctl/p5r-stats/src/services"
)

func UserFormHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/partials/user/edit/"):]

	switch r.Method {

	case http.MethodGet:
		render.HTML(w, render.FragmentUserEdit, models.Username{Name: name}, nil)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func UserInfoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/partials/user/edit-cancel/"):]

	switch r.Method {

	case http.MethodGet:
		render.HTML(w, render.FragmentUsernameDiv, models.Username{Name: name}, nil)

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

		render.HTML(w, render.FragmentHome, data, []render.OOB{
			{
				ID:   "toast-container",
				Swap: "beforeend",
				View: render.FragmentToast,
				Data: models.Toast{
					Type:    "success",
					Message: "User created",
				},
			},
		})
	
	case http.MethodPatch:
		fmt.Printf("r.Header.Get(\"Content-Type\"): %v\n", r.Header.Get("Content-Type"))

		// 10 MB max
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			http.Error(w, "Bad Form", http.StatusBadRequest)
			return
		}

		pfp, ok := r.MultipartForm.File["pfp"]
		fmt.Println("pfp:", ok)

		// name := r.Form.Get("name")
		name, ok := r.MultipartForm.Value["name"]
		fmt.Println("name:", ok)

		fmt.Println("pfp:", pfp)
		fmt.Println("name:", name)

		fmt.Printf("name[0]: %v\n", name[0])

		fmt.Printf("len(name): %v\n", len(name))
		if len(name[0]) == 0 {
			http.Error(w, "Name is required.", http.StatusBadRequest)
			return
		}

		render.HTML(w, render.FragmentUsernameDiv, models.Username{Name: name[0]}, nil)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func ModalHandler(w http.ResponseWriter, r *http.Request) {
	// content := r.URL.Path[len("/partials/modal/"):]

	switch r.Method {
	case http.MethodGet:
		data := models.Modal{
			Title:   "Add Activity",
			Content: "activity",
		}
		render.HTML(w, render.FragmentModal, data, nil)
	}
}
