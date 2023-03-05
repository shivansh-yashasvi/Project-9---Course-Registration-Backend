package structs

import "time"

type TimeTable struct {
	Success bool          `json:"success"`
	Data    TimeTableBody `json:"data"`
}

type TimeTableBody struct {
	StudentID         string                 `json:"id"`
	StudentName       string                 `json:"name"`
	RegisteredCourses []RegisteredCourseBody `json:"registered_courses"`
}

type RegisteredCourseBody struct {
	Course CourseBody  `json:"course"`
	Slots  []SlotsBody `json:"slots"`
}

type CourseBody struct {
	CourseID      string              `json:"id"`
	Name          string              `json:"name"`
	Faculties     []FacultiesBody     `json:"faculties"`
	CourseType    string              `json:"course_type"`
	AllowedSlotes []AllowedSlotesBody `json:"allowed_slots"`
}

type FacultiesBody struct {
	FacultyID string `json:"id"`
	Name      string `json:"name"`
}

type AllowedSlotesBody struct {
	AllowedSlotesID string                      `json:"id"`
	Timings         []AllowedSlotesBodyTimeBody `json:"timings"`
}

type AllowedSlotesBodyTimeBody struct {
	Day   string    `json:"day"`
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

type SlotsBody struct {
	SlotesID string                      `json:"id"`
	Timings  []AllowedSlotesBodyTimeBody `json:"timings"`
}

type TimeTableRepo struct {
	Success          bool
	StudentID        string
	StudentName      string
	CourseID         string
	CourseName       string
	CourseType       string
	FacultyID        string
	SlotesID         string
	Day              string
	Start            time.Time
	End              time.Time
	FacultyName      string
	AllowedSlotesID  string
	AllowedSlotDay   string
	AllowedSlotStart time.Time
	AllowedSlotEnd   time.Time
}
