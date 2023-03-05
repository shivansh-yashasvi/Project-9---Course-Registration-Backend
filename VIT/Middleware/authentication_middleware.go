package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
)

var secretkey = "super-secret-key"

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] == nil {
			err := errors.New("no token found")
			w.WriteHeader(412)
			json.NewEncoder(w).Encode(err)
			return
		}

		var mySigningKey = []byte(secretkey)

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing")
			}
			return mySigningKey, nil
		})

		if err != nil {
			err := errors.New("token is expired")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "admin" {

				r.Header.Set("Role", "admin")
				handler.ServeHTTP(w, r)
				return

			} else {
				if claims["role"] == "student" {

					r.Header.Set("Role", "student")
					handler.ServeHTTP(w, r)
					return

				}
			}
		}
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func IsAdminAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] == nil {
			err := errors.New("no token found")
			w.WriteHeader(412)
			json.NewEncoder(w).Encode(err)
			return
		}

		var mySigningKey = []byte(secretkey)

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing")
			}
			return mySigningKey, nil
		})

		if err != nil {
			err := errors.New("token is expired")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "admin" {

				r.Header.Set("Role", "admin")
				handler.ServeHTTP(w, r)
				return

			}
		}
		w.WriteHeader(http.StatusUnauthorized)
	}
}
func IsStudentAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] == nil {
			err := errors.New("no token found")
			w.WriteHeader(412)
			json.NewEncoder(w).Encode(err)
			return
		}

		var mySigningKey = []byte(secretkey)

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing")
			}
			return mySigningKey, nil
		})

		if err != nil {
			err := errors.New("token is expired")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "student" {

				r.Header.Set("Role", "student")
				if val, ok := claims["ID"]; ok {
					r.Header.Set("ID", fmt.Sprintf("%v", val))
				}

				handler.ServeHTTP(w, r)
				return

			}
		}
		w.WriteHeader(http.StatusUnauthorized)
	}
}

// func IsStudentAuthorized(handler http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		if r.Header["Token"] == nil {
// 			err := errors.New("no token found")
// 			w.WriteHeader(412)
// 			json.NewEncoder(w).Encode(err)
// 			return
// 		}

// 		var mySigningKey = []byte(secretkey)

// 		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("There was an error in parsing")
// 			}
// 			return mySigningKey, nil
// 		})

// 		if err != nil {
// 			err := errors.New("token is expired")
// 			w.WriteHeader(http.StatusUnauthorized)
// 			json.NewEncoder(w).Encode(err)
// 			return
// 		}

// 		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 			if claims["role"] == "student" {

// 				r.Header.Set("Role", "student")
// 				handler.ServeHTTP(w, r)
// 				return

// 			}
// 		}
// 		w.WriteHeader(http.StatusUnauthorized)
// 	}
// }
