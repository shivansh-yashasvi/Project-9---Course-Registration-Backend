package services

import (
	"VIT/repository"
	"VIT/structs"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func SignUp(user structs.User) (userID string, err error) {
	exist, err := repository.CheckEmail(user.Email)
	if err != nil {
		return userID, nil
	}
	if exist {
		return userID, errors.New("user already exist")
	}

	user.Password, err = GeneratehashPassword(user.Password)
	if err != nil {
		log.Fatalln("error in password hash")
	}

	userID, err = repository.SignUp(user)
	if err != nil {
		return userID, nil
	}

	return userID, nil
}

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
