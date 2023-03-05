package repository

import (
	"VIT/DBConnection"
	"VIT/structs"
	"encoding/json"

	"github.com/google/uuid"
)

func CreateRegister(register structs.Register) (structs.Register, error) {
	id := uuid.New().String()

	out, err := json.Marshal(register.SlotID)
	if err != nil {
		return structs.Register{}, err
	}
	strSlotIDs := string(out)

	dbm := DBConnection.ConDB()
	_, err = dbm.Exec("insert into register (id,course_id,faculty_id,slot_id,student_id) values(?,?,?,?,?)", id, register.CourseID, register.FacultyID, strSlotIDs, register.StudentID)
	if err != nil {
		return register, err
	}
	return register, nil
}
func GetRegisterByID(studentID, courseID string) (structs.Register, error) {
	var register structs.Register
	dbm := DBConnection.ConDB()
	rows, err := dbm.Query("select student_id,course_id,faculty_id,slot_id from register where course_id = ? and student_id = ?", courseID, studentID)
	if err != nil {
		return register, err
	}
	for rows.Next() {
		rows.Scan(&register.StudentID, &register.CourseID, &register.FacultyID, &register.SlotID)
		if err != nil {
			return register, err
		}
	}
	return register, nil
}
