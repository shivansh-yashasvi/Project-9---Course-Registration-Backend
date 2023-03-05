package repository

import (
	"VIT/DBConnection"
	"VIT/structs"
	"fmt"
)

func CreateFaculty(faculty structs.Faculty) (structs.Faculty, error) {

	dbm := DBConnection.ConDB()
	_, err := dbm.Exec("insert ignore into faculty (id,name) values(?,?)", faculty.ID, faculty.Name)
	if err != nil {
		return faculty, err
	}
	return faculty, nil
}
func GetFacultyByID(ID string) (structs.Faculty, error) {
	var faculty structs.Faculty
	dbm := DBConnection.ConDB()
	fmt.Println(ID)
	rows, err := dbm.Query("select * from faculty where id=?", ID)
	if err != nil {
		return structs.Faculty{}, err
	}
	for rows.Next() {
		err := rows.Scan(&faculty.ID, &faculty.Name)
		if err != nil {
			return structs.Faculty{}, err
		}

	}
	return faculty, nil
}
