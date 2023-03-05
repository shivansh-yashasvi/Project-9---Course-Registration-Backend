package services

import (
	"VIT/repository"
	"VIT/structs"
)

func CreateSlot(slot structs.Slot) (structs.Data, error) {
	return repository.CreateSlot(slot)
}
