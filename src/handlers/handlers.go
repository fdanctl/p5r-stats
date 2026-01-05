package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/fdanctl/p5r-stats/src/models"
	"github.com/fdanctl/p5r-stats/src/render"
	"github.com/fdanctl/p5r-stats/src/services"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		userData, err := services.ReadUserData()

		if err != nil {
			render.HTML(w, "new.html", nil)
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
			render.HTML(w, "home.html", data)
		}

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func UserDataHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodPost:
		err := r.ParseForm()

		if err != nil {
			http.Error(w, "Method Not Allowed", http.StatusBadRequest)
			return
		}

		err = services.CreateUserData(r.Form.Get("name"))
		if err != nil {
			if errors.Is(err, models.ErrAlreadyExists) {
				http.Error(w, "User data found", http.StatusBadRequest)
				return
			}
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)

	case http.MethodPatch:
		fmt.Println("modify user data. soon")

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

func DesignHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		render.HTML(w, "design-system.html", nil)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func ActivityHandler(w http.ResponseWriter, r *http.Request) {
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

		decoder := json.NewDecoder(r.Body)
		var body models.ActivityInput
		err := decoder.Decode(&body)
		if err != nil {
			var syntaxErr *json.SyntaxError
			if errors.As(err, &syntaxErr) {
				http.Error(
					w,
					fmt.Sprintf("Malformed JSON at byte %d", syntaxErr.Offset),
					http.StatusBadRequest,
				)
				return
			}

			var unmarshalTypeErr *json.UnmarshalTypeError
			if errors.As(err, &unmarshalTypeErr) {
				http.Error(
					w,
					fmt.Sprintf("Field '%s' has wrong type", unmarshalTypeErr.Field),
					http.StatusBadRequest,
				)
				return
			}

			if errors.Is(err, models.ErrInvalidStat) {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			http.Error(w, "Invalid request body", http.StatusBadRequest)
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

func ActivityIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/activity/"):]

	switch r.Method {
	case http.MethodPost:
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(
				w,
				"Content-Type must be application/json",
				http.StatusUnsupportedMediaType,
			)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var body models.ActivityModifyInput
		err := decoder.Decode(&body)
		if err != nil {
			var syntaxErr *json.SyntaxError
			if errors.As(err, &syntaxErr) {
				http.Error(
					w,
					fmt.Sprintf("Malformed JSON at byte %d", syntaxErr.Offset),
					http.StatusBadRequest,
				)
				return
			}

			var unmarshalTypeErr *json.UnmarshalTypeError
			if errors.As(err, &unmarshalTypeErr) {
				http.Error(
					w,
					fmt.Sprintf("Field '%s' has wrong type", unmarshalTypeErr.Field),
					http.StatusBadRequest,
				)
				return
			}

			if errors.Is(err, models.ErrInvalidStat) {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
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
