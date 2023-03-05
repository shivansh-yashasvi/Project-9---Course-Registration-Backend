package controllers

import (
	"VIT/controllers/validation"
	"VIT/services"
	"VIT/structs"
	"encoding/json"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user structs.User
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	validErr := validation.Struct(user)
	if validErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userID, err := services.SignUp(user)
	if err != nil {
		if err.Error() == "user already exist" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(userID)
}
