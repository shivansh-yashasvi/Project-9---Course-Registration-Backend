package controllers

import (
	"VIT/controllers/validation"
	"VIT/services"
	"VIT/structs"
	"encoding/json"
	"net/http"
)

func CreateRegister(w http.ResponseWriter, r *http.Request) {
	var register structs.Register
	ID := r.Header.Get("ID")
	if ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	register.StudentID = ID

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&register)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	validErr := validation.Struct(register)
	if validErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	register, err = services.CreateRegister(register)
	if err != nil {
		if err.Error() == "faculty not exist for that course" {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		if err.Error() == "slot not exist for that course" {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		if err.Error() == "student is already registered for the given course" {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		if err.Error() == "the choosen slot timings are already registered" {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(register)

}
