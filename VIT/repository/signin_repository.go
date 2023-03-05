package repository

import (
	"VIT/DBConnection"
	"VIT/structs"
)

func GetUserByEmail(email string) (structs.User, bool, error) {
	var user structs.User
	exist := false
	dbm := DBConnection.ConDB()
	rows, err := dbm.Query("select email,hashPassword,role from admin where email = ?", email)
	if err != nil {
		return structs.User{}, exist, nil
	}
	for rows.Next() {
		rows.Scan(&user.Email, &user.Password, &user.Role)
		if err != nil {
			return structs.User{}, exist, nil
		}
		exist = true
	}
	return user, exist, nil
}
