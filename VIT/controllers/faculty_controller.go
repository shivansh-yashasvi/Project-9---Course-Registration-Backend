package controllers

import (
	"VIT/controllers/validation"
	"VIT/services"
	"VIT/structs"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type FacultyData struct {
	Success bool            `json:"success"`
	Data    structs.Faculty `json:"data"`
}

func CreateFaculty(w http.ResponseWriter, r *http.Request) {
	var faculty structs.Faculty
	var facultydata FacultyData
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&faculty)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		facultydata.Success = false
		json.NewEncoder(w).Encode(facultydata)
		return
	}

	validErr := validation.Struct(faculty)
	if validErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		facultydata.Success = false
		json.NewEncoder(w).Encode(facultydata)
		return
	}
	faculty, err = services.CreateFaculty(faculty)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		facultydata.Success = false
		json.NewEncoder(w).Encode(facultydata)
		return
	}
	facultydata.Success = true
	facultydata.Data = faculty

	json.NewEncoder(w).Encode(facultydata)

}

func GetFacultyByID(w http.ResponseWriter, r *http.Request) {
	var facultydata FacultyData
	vars := mux.Vars(r)
	ID := vars["faculty_id"]
	if ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		facultydata.Success = false
		json.NewEncoder(w).Encode(facultydata)
		return
	}
	faculty, err := services.GetFacultyByID(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		facultydata.Success = false
		json.NewEncoder(w).Encode(facultydata)
		return
	}
	if faculty.ID == "" && faculty.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		facultydata.Success = false
		json.NewEncoder(w).Encode(facultydata)
		return
	}
	facultydata.Success = true
	facultydata.Data = faculty
	json.NewEncoder(w).Encode(facultydata)

}
