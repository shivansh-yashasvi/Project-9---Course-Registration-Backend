package repository

import (
	"VIT/DBConnection"
	"VIT/structs"
	"encoding/json"
	"fmt"
)

func GetTimeTable(StudentID string) (structs.TimeTableRepo, []string, error) {
	var timetable structs.TimeTableRepo
	var slotIDArr []string
	dbm := DBConnection.ConDB()
	query := `
			select 
				r.student_id,
				s.name,
				r.course_id,
				c.name,
				c.course_type,
				r.faculty_id,
				f.name,
				r.slot_id
				from register r
			INNER JOIN student s 
				ON r.student_id = s.id
			INNER JOIN course c 
				ON r.course_id = c.id
			INNER JOIN faculty f 
				ON r.faculty_id = f.id
	 			where student_id = ?`
	rows, err := dbm.Query(query, StudentID)
	if err != nil {
		return timetable, slotIDArr, err
	}
	for rows.Next() {
		var studentID, studentName, courseID, courseName, courseType, facultyID, facultyName, slotID string
		rows.Scan(&studentID, &studentName, &courseID, &courseName, &courseType, &facultyID, &facultyName, &slotID)
		if err != nil {
			return timetable, slotIDArr, err
		}

		err = json.Unmarshal([]byte(slotID), &slotIDArr)
		if err != nil {
			return timetable, slotIDArr, err
		}
		fmt.Println(slotIDArr)
		// slotIDArr = strings.Split(strings.Trim(slotIDArr[0], "\""), ",")

		timetable = structs.TimeTableRepo{
			Success:     true,
			StudentID:   studentID,
			StudentName: studentName,
			CourseID:    courseID,
			CourseName:  courseName,
			CourseType:  courseType,
			FacultyID:   facultyID,
			FacultyName: facultyName,
		}
	}
	return timetable, slotIDArr, nil
}

// func GetTimeTable(StudentID string) ([]structs.TimeTable, error) {
// 	var timetable []structs.TimeTable
// 	dbm := DBConnection.ConDB()
// 	rows, err := dbm.Query("select * from timetable where student_id = ?", StudentID)
// 	if err != nil {
// 		return timetable, err
// 	}
// 	for rows.Next() {
// 		rows.Scan(&timetable.StudentID, &timetable.CourseID, &timetable.FacultyID, &timetable.SlotID)
// 		if err != nil {
// 			return timetable, nil
// 		}
// 	}
// 	return timetable, nil
// }
