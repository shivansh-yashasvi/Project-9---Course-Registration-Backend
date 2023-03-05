package services

import (
	"VIT/repository"
	"VIT/structs"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var secretkey = "super-secret-key"

func SignIn(auth structs.Authentication) (structs.Token, error) {
	var token structs.Token
	authUser, exist, err := repository.GetUserByEmail(auth.Email)
	if err != nil {
		return token, nil
	}
	if !exist {
		return token, errors.New("email or password is incorrect")
	}

	check := CheckPasswordHash(auth.Password, authUser.Password)

	if !check {
		return token, errors.New("email or password is incorrect")
	}

	validToken, err := GenerateJWT("", authUser.Email, "admin")
	if err != nil {
		return token, err
	}

	token.Email = authUser.Email
	token.Role = authUser.Role
	token.TokenString = validToken

	return token, nil
}

func StudentSignIn(body structs.StudentSignInBody) (structs.StudentToken, error) {
	var token structs.StudentToken
	ID, _, hashPassword, exist, err := repository.GetStudentByID(body.ID)
	if err != nil {
		return token, nil
	}
	if !exist {
		return token, errors.New("ID or password is incorrect")
	}

	check := CheckPasswordHash(body.Password, hashPassword)

	if !check {
		return token, errors.New("ID or password is incorrect")
	}

	validToken, err := GenerateJWT(ID, "", "student")
	if err != nil {
		return token, err
	}

	token.ID = ID
	token.Role = body.Role
	token.TokenString = validToken

	return token, nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT(ID, email, role string) (string, error) {
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["ID"] = ID
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
