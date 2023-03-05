package services

import (
	"VIT/repository"
	"VIT/structs"
)

func CreateFaculty(faculty structs.Faculty) (structs.Faculty, error) {
	return repository.CreateFaculty(faculty)
}

func GetFacultyByID(faculty string) (structs.Faculty, error) {
	return repository.GetFacultyByID(faculty)
}
