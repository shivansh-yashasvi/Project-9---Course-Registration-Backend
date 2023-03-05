package structs

type Register struct {
	StudentID string   `json:"student_id"`
	CourseID  string   `json:"course_id" validate:"required"`
	FacultyID string   `json:"faculty_id" validate:"required"`
	SlotID    []string `json:"slot_id" validate:"required"`
}
