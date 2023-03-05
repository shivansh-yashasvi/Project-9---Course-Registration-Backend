package controllers

import (
	"VIT/controllers/validation"
	"VIT/services"
	"VIT/structs"
	"encoding/json"
	"net/http"
)

type Studentdata struct {
	Success bool                    `json:"success"`
	Data    structs.StudentResponse `json:"data"`
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student structs.Student
	var data Studentdata
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&student)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data.Success = false
		json.NewEncoder(w).Encode(data)
		return
	}

	validErr := validation.Struct(student)
	if validErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		data.Success = false
		json.NewEncoder(w).Encode(data)
		return
	}
	dataresponse, err := services.CreateStudent(student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		data.Success = false
		json.NewEncoder(w).Encode(data)
		return
	}
	data.Success = true
	data.Data.ID = dataresponse.ID
	data.Data.Name = dataresponse.Name
	json.NewEncoder(w).Encode(data)

}
