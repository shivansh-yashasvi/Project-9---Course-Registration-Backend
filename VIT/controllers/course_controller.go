package controllers

import (
	"VIT/controllers/validation"
	"VIT/services"
	"VIT/structs"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	var course structs.Course
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&course)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	validErr := validation.Struct(course)
	if validErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	courseData, err := services.CreateCourse(course)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(courseData)

}

func GetCourseByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["course_id"]
	if ID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	course, err := services.GetCourseByID(ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(course)

}
