package handlers

import (
	"fmt"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "ok")

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
