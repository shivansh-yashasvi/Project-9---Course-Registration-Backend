package repository

import (
	"VIT/DBConnection"
	"VIT/structs"

	"github.com/google/uuid"
)

func SignUp(user structs.User) (string, error) {
	//generates new uuid
	uuid := uuid.New()
	id := uuid.String()

	dbm := DBConnection.ConDB()
	_, err := dbm.Exec("insert into admin (id,name,email,hashPassword,role) values(?,?,?,?,?)", id, user.Name, user.Email, user.Password, "admin")
	if err != nil {
		return id, err
	}
	return id, nil
}

func CheckEmail(email string) (bool, error) {
	exist := false
	dbm := DBConnection.ConDB()
	rows, err := dbm.Query("select email from admin where email = ?", email)
	if err != nil {
		return exist, err
	}
	for rows.Next() {
		rows.Scan(&email)
		if err != nil {
			return exist, nil
		}
		exist = true
	}
	return exist, nil
}
