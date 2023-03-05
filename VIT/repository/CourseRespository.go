package repository

import (
	"VIT/DBConnection"
	"VIT/structs"
	"encoding/json"
	"strings"
)

func CreateCourse(course structs.Course) (courseData structs.CourseData, err error) {
	out, err := json.Marshal(course.FacultyIDs)
	if err != nil {
		return structs.CourseData{}, nil
	}
	facultyIDs := string(out)
	out2, err := json.Marshal(course.SlotIDs)
	if err != nil {
		return structs.CourseData{}, nil
	}
	slotIDs := string(out2)

	dbm := DBConnection.ConDB()
	_, err = dbm.Exec("insert ignore into course (id,name,faculty_id,course_type,slot_id) values(?,?,?,?,?)", course.ID, course.Name, facultyIDs, course.CourseType, slotIDs)
	if err != nil {
		return structs.CourseData{}, err
	}

	facultyObj := make([]structs.Faculty, 0)

	for _, fID := range course.FacultyIDs {
		faculty, err := GetFacultyByID(fID)
		if err != nil {
			return structs.CourseData{}, err
		}
		facultyObj = append(facultyObj, structs.Faculty{
			ID:   faculty.ID,
			Name: faculty.Name,
		})
	}

	slotObj := make([]structs.Slot, 0)

	for _, sID := range course.SlotIDs {
		slot, err := GetSlotByID(sID)
		if err != nil {
			return structs.CourseData{}, err
		}
		slotObj = append(slotObj, structs.Slot{
			ID:    slot.ID,
			Day:   slot.Day,
			Start: slot.Start,
			End:   slot.End,
		})
	}

	courseData = structs.CourseData{
		Success: false,
		Data: structs.CourseResponse{
			ID:         course.ID,
			Name:       course.Name,
			FacultyIDs: facultyObj,
			SlotIDs:    slotObj,
			CourseType: course.CourseType,
		},
	}

	return courseData, nil
}
func GetCourseByID(ID string) (structs.Course, error) {
	var course structs.Course
	var facultyIDStr, slotIDStr string
	var facultyIDArr, slotIDArr []string
	dbm := DBConnection.ConDB()
	rows, err := dbm.Query("select * from course where id=?", ID)
	for rows.Next() {
		rows.Scan(&course.ID, &course.Name, &facultyIDStr, &course.CourseType, &slotIDStr)
		if err != nil {
			return course, nil
		}
		if err != nil {
			return course, err
		}

	}
	err = json.Unmarshal([]byte(facultyIDStr), &facultyIDArr)
	if err != nil {
		return course, err
	}
	facultyIDArr = strings.Split(strings.Trim(facultyIDArr[0], "\""), ",")
	course.FacultyIDs = facultyIDArr

	err = json.Unmarshal([]byte(slotIDStr), &slotIDArr)
	if err != nil {
		return course, err
	}
	slotIDArr = strings.Split(strings.Trim(slotIDArr[0], "\""), ",")
	course.SlotIDs = slotIDArr

	return course, nil
}
