package handlers

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"

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

		// name := r.Form.Get("name")
		name, ok := r.MultipartForm.Value["name"]
		if !ok || len(name[0]) == 0 {
			http.Error(w, "Name is required.", http.StatusBadRequest)
			return
		}

		var fh *multipart.FileHeader
		pfp, ok := r.MultipartForm.File["pfp"]
		if ok && len(pfp) > 0 {
			fh = pfp[0]
		}

		err := services.ModifyUser(name[0], fh)
		if err != nil {
			errors.Is(err, models.ErrCantReadFile)
			http.Error(w, "Can't read the file", http.StatusBadRequest)
		} else {
			http.Error(w, "Failed to save file", http.StatusInternalServerError)
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

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func ActivityHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad Form", http.StatusBadRequest)
			return
		}
		title := r.Form.Get("title")
		if title == "" {
			http.Error(w, "Name is required.", http.StatusBadRequest)
			return
		}

		points := r.Form["points"]
		stats := r.Form["stat"]
		fmt.Printf("stats: %v\n", stats)
		if len(stats) == 0 || len(stats) > len(points) {
			http.Error(w, "At least one stat must be provided", http.StatusBadRequest)
			return
		}

		is := make([]models.IncreasedStat, 0, 5)
		for i, s := range stats {
			stat, err := models.ParseStat(s)
			if err != nil {
				http.Error(
					w,
					"Invalid stat\nStats must be Knowledge, Guts, Proficiency, Kindness or charm",
					http.StatusBadRequest,
				)
				return
			}

			p, err := strconv.ParseUint(points[i], 0, 8)
			if err != nil {
				http.Error(w, "Invalid number in files points", http.StatusBadRequest)
				return
			}
			fmt.Printf("p: %v\n", p)
			if p <= 0 || p > 10 {
				http.Error(
					w,
					"Invalid number range. Increased points must be between 0 and 10",
					http.StatusBadRequest,
				)
				return
			}

			is = append(is, models.IncreasedStat{Stat: stat, Points: uint8(p)})

		}

		aData, err := services.InsertActivity(models.ActivityInput{
			Title:          title,
			Description:    r.Form.Get("description"),
			IncreasedStats: is,
		})

		if err != nil {
			http.Error(w, "Bad Form", http.StatusBadRequest)
			return
		}

		userData, err := services.ReadUserData()
		statsMap := services.ComputeStats(userData.Activities)

		gData := models.Stats{
			Knowledge:   statsMap[models.Knowledge],
			Guts:        statsMap[models.Guts],
			Proficiency: statsMap[models.Proficiency],
			Kindness:    statsMap[models.Kindness],
			Charm:       statsMap[models.Charm],
		}

		render.RenderOOB(w, "activity-list", "afterbegin", render.FragmentActivity, aData)
		render.RenderOOB(w, "stats-graph", "outerHTML", render.FragmentStatsGraph, gData)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func StatHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		queryStat := r.URL.Query()["stat"]
		if len(queryStat) >= 5 {
			http.StatusText(http.StatusNoContent)
			return
		}

		options := []models.Stat{
			models.Knowledge,
			models.Guts,
			models.Proficiency,
			models.Kindness,
			models.Charm,
		}
		for _, v := range queryStat {
			s, _ := models.ParseStat(v)

			for i, o := range options {
				if o == s {
					options = append(options[0:i], options[i+1:]...)
					break
				}
			}
		}

		type optionsStruct struct {
			Options []string
		}

		opts := make([]string, 0, 5)
		for _, v := range options {
			opts = append(opts, v.String())
		}

		fmt.Printf("options: %v\n", options)
		render.HTML(w, render.FragmentStatSelect, optionsStruct{
			Options: opts,
		}, nil)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

	}
}
