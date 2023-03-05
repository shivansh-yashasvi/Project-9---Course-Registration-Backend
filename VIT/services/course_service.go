package services

import (
	"VIT/repository"
	"VIT/structs"
)

func CreateCourse(course structs.Course) (structs.CourseData, error) {
	coursedata, err := repository.CreateCourse(course)
	if err != nil {
		return structs.CourseData{}, err
	}
	return coursedata, nil
}

func GetCourseByID(course string) (structs.Course, error) {
	return repository.GetCourseByID(course)
}
