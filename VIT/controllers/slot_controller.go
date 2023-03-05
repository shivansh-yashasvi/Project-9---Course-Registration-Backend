package controllers

import (
	"VIT/controllers/validation"
	"VIT/services"
	"VIT/structs"
	"encoding/json"
	"net/http"
)

func CreateSlot(w http.ResponseWriter, r *http.Request) {
	var slot structs.Slot
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&slot)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	validErr := validation.Struct(slot)
	if validErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data, err := services.CreateSlot(slot)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(data)

}
