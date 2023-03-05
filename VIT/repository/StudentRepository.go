package repository

import (
	"VIT/DBConnection"
	"VIT/structs"
)

func CreateStudent(ID string, student structs.Student) (structs.Student, error) {
	dbm := DBConnection.ConDB()
	student.ID = ID
	_, err := dbm.Exec("insert ignore into student (id,name,hashPassword) values(?,?,?)", student.ID, student.Name, student.Password)
	if err != nil {
		return student, err
	}

	return student, nil
}

func GetStudentCount() (int64, error) {
	var studentCount int64
	dbm := DBConnection.ConDB()
	rows, err := dbm.Query("select count(id) from student")
	if err != nil {
		return studentCount, err
	}
	for rows.Next() {
		rows.Scan(&studentCount)
		if err != nil {
			return studentCount, err
		}
	}
	return studentCount, nil
}
func GetStudentByID(ID string) (string, string, string, bool, error) {
	var id, name, hashPassword string
	exist := false
	dbm := DBConnection.ConDB()
	rows, err := dbm.Query("select id,name,hashPassword from student where id = ?", ID)
	if err != nil {
		return id, name, hashPassword, exist, err
	}
	for rows.Next() {
		rows.Scan(&id, &name, &hashPassword)
		if err != nil {
			return id, name, hashPassword, exist, err
		}
		exist = true
	}
	return id, name, hashPassword, exist, nil
}
