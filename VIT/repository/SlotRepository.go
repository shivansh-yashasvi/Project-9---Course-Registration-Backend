package repository

import (
	"VIT/DBConnection"
	"VIT/structs"
)

func CreateSlot(slot structs.Slot) (data structs.Data, err error) {

	dbm := DBConnection.ConDB()
	_, err = dbm.Exec("insert ignore into slot (id,day,start,end) values(?,?,?,?)", slot.ID, slot.Day, slot.Start, slot.End)
	if err != nil {
		return data, err
	}
	data = structs.Data{
		Success: true,
		Data: structs.Slot{
			ID:    slot.ID,
			Day:   slot.Day,
			Start: slot.Start,
			End:   slot.End,
		},
	}
	return data, nil
}
func GetSlotByID(slotID string) (structs.Slot, error) {
	var slot structs.Slot
	dbm := DBConnection.ConDB()
	rows, err := dbm.Query("select id,day,start,end from slot where id = ?", slotID)
	if err != nil {
		return slot, err
	}
	for rows.Next() {
		rows.Scan(&slot.ID, &slot.Day, &slot.Start, &slot.End)
		if err != nil {
			return slot, err
		}
	}
	return slot, nil
}
