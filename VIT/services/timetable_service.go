package services

import (
	"VIT/repository"
	"VIT/structs"
)

var allowedSlotBody = make([]structs.AllowedSlotesBody, 0)

var SlotsBody = make([]structs.SlotsBody, 0)

func GetTimeTable(ID string) (timeTable structs.TimeTable, err error) {
	timetableRepo, slotIdArr, err := repository.GetTimeTable(ID)
	if err != nil {
		return timeTable, nil
	}

	for _, slotID := range slotIdArr {
		slot, err := repository.GetSlotByID(slotID)
		if err != nil {
			return timeTable, err
		}
		allowedSlotBody = append(allowedSlotBody, structs.AllowedSlotesBody{
			AllowedSlotesID: slot.ID,
			Timings: []structs.AllowedSlotesBodyTimeBody{
				structs.AllowedSlotesBodyTimeBody{
					Day:   slot.Day,
					Start: slot.Start,
					End:   slot.End,
				},
			},
		})
		SlotsBody = append(SlotsBody, structs.SlotsBody{
			SlotesID: slot.ID,
			Timings: []structs.AllowedSlotesBodyTimeBody{
				structs.AllowedSlotesBodyTimeBody{
					Day:   slot.Day,
					Start: slot.Start,
					End:   slot.End,
				},
			},
		})
	}

	timeTable = structs.TimeTable{
		Success: timetableRepo.Success,
		Data: structs.TimeTableBody{
			StudentID:   timetableRepo.StudentID,
			StudentName: timetableRepo.StudentName,
			RegisteredCourses: []structs.RegisteredCourseBody{
				structs.RegisteredCourseBody{
					Course: structs.CourseBody{
						CourseID:   timetableRepo.CourseID,
						Name:       timetableRepo.StudentName,
						CourseType: timetableRepo.CourseType,
						Faculties: []structs.FacultiesBody{
							structs.FacultiesBody{
								FacultyID: timetableRepo.FacultyID,
								Name:      timetableRepo.FacultyName,
							},
						},
						AllowedSlotes: allowedSlotBody,
					},
					Slots: SlotsBody,
				},
			},
		},
	}

	return timeTable, nil
}
