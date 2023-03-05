package controllers

import (
	"VIT/services"
	"encoding/json"
	"net/http"
)

func GetTimeTable(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("ID")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	timetable, err := services.GetTimeTable(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(timetable)
}
