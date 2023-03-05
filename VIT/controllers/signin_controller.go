package controllers

import (
	"VIT/controllers/validation"
	"VIT/services"
	"VIT/structs"
	"encoding/json"
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	var user structs.Authentication
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

	token, err := services.SignIn(user)
	if err != nil {
		if err.Error() == "email or password is incorrect" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(token)
}

func StudentSignIn(w http.ResponseWriter, r *http.Request) {
	var user structs.StudentSignInBody
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

	token, err := services.StudentSignIn(user)
	if err != nil {
		if err.Error() == "ID or password is incorrect" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(token)
}
