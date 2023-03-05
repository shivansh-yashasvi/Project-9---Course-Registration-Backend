package structs

type Course struct {
	ID         string   `json:"id"         validate:"required"`
	Name       string   `json:"name"       validate:"required"`
	FacultyIDs []string `json:"faculty_id"`
	SlotIDs    []string `json:"slot_id"`
	CourseType string   `json:"course_type" validate:"required,oneof=THEORY LAB"`
}

type CourseResponse struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	FacultyIDs []Faculty `json:"faculties"`
	SlotIDs    []Slot    `json:"allowed_slots"`
	CourseType string    `json:"course_type"`
}

type CourseData struct {
	Success bool           `json:"success"`
	Data    CourseResponse `json:"data"`
}
