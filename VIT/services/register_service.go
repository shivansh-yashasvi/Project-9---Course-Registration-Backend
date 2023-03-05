package services

import (
	"VIT/repository"
	"VIT/structs"
	"errors"
)

func CreateRegister(body structs.Register) (structs.Register, error) {

	registeredStudent, err := repository.GetRegisterByID(body.StudentID, body.CourseID)
	if err != nil {
		return structs.Register{}, err
	}

	//check if student is already registered for given courseID
	if body.CourseID == registeredStudent.CourseID {
		return structs.Register{}, errors.New("student is already registered for the given course")
	}

	faculty, err := repository.GetCourseByID(body.CourseID)
	if err != nil {
		return structs.Register{}, err
	}

	if len(faculty.FacultyIDs) == 0 {
		return structs.Register{}, errors.New("faculty not exist for that course")
	}

	//check if the chooesen faculty should teach that course
	for _, id := range faculty.FacultyIDs {

		if body.FacultyID == id {
			break
		} else {
			return structs.Register{}, errors.New("faculty not exist for that course")
		}
	}

	// check if chosen slots should exist for that course
	if len(faculty.SlotIDs) == 0 {
		return structs.Register{}, errors.New("slot not exist for that course")
	}

	for _, slotID := range body.SlotID {
		exist := arrContains(faculty.SlotIDs, slotID)
		if exist {
			continue
		} else {
			return structs.Register{}, errors.New("slot not exist for that course")
		}
	}

	// The chosen slot for one course should not clash with a different registered course's slot.
	for _, slotID := range body.SlotID {
		requestedSlot, err := repository.GetSlotByID(slotID)
		if err != nil {
			return structs.Register{}, err
		}

		for _, id := range registeredStudent.SlotID {
			registeredSlot, err := repository.GetSlotByID(id)
			if err != nil {
				return structs.Register{}, err
			}

			if requestedSlot.Day == registeredSlot.Day && requestedSlot.Start == registeredSlot.Start && requestedSlot.End == registeredSlot.End {
				return structs.Register{}, errors.New("the choosen slot timings are already registered")
			}
		}

	}

	return repository.CreateRegister(body)
}

func arrContains(a []string, b string) bool {
	for _, v := range a {
		if b == v {
			return true
		}
	}
	return false
}
