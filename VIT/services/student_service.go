package services

import (
	"VIT/repository"
	"VIT/structs"
	"fmt"
	"log"
	"strconv"
)

func CreateStudent(student structs.Student) (structs.Student, error) {
	var ID string
	count, err := repository.GetStudentCount()
	if err != nil {
		return student, err
	}
	count = count + 1
	ID = fmt.Sprintf("2OBCE60"+"%s", strconv.FormatInt(count, 10))
	student.Password, err = GeneratehashPassword(student.Password)
	if err != nil {
		log.Fatalln("error in password hash")
	}
	return repository.CreateStudent(ID, student)
}
