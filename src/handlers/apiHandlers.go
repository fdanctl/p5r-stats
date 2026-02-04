package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/fdanctl/p5r-stats/src/middleware"
	"github.com/fdanctl/p5r-stats/src/models"
	"github.com/fdanctl/p5r-stats/src/services"
)

func UserDataHandlerAPI(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodPatch:
		// fmt.Println("modify user data. soon")
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

	case http.MethodDelete:
		err := services.DeleteUserData()
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func ActivityHandlerAPI(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodPost:
		// TODO:
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(
				w,
				"Content-Type must be application/json",
				http.StatusUnsupportedMediaType,
			)
			return
		}

		body, err := middleware.DecodeRequestBody[models.ActivityInput](r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = body.Validate()
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusBadRequest)
			return
		}

		err = services.InsertActivity(body)
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func ActivityIdHandlerAPI(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/api/activity/"):]

	switch r.Method {
	case http.MethodGet:
		fmt.Println("getting activity with id:", id)
		activ, err := services.ReadActivity(id)
		fmt.Println(activ)
		if err != nil {
			if errors.Is(err, models.ErrNotFound) {
				http.Error(w, "Activity Not Found", http.StatusBadRequest)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(activ); err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
		}

	case http.MethodPost:
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(
				w,
				"Content-Type must be application/json",
				http.StatusUnsupportedMediaType,
			)
			return
		}

		body, err := middleware.DecodeRequestBody[models.ActivityModifyInput](r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = body.Validate()
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusBadRequest)
			return
		}

		err = services.ModifyActivity(id, body)
		fmt.Printf("err: %v\n", err)

	case http.MethodDelete:
		err := services.DeleteActivity(id)
		// TODO: handle errors
		if err != nil {
			if errors.Is(err, models.ErrNotFound) {
				http.Error(w, "Activity not found", http.StatusBadRequest)
			}
			http.Error(w, "Internal error", http.StatusInternalServerError)
		}
		// TODO: response
	}
}
