package services

import (
	"backend/domain"
	"time"
)

func GetActivityByID(id int) domain.Activity {
	return domain.Activity{
		ID:          1,
		Name:        "Zumba",
		Description: "Clase de Zumba",
		Schedule: []domain.TimeSlot{
			{
				ID:         1,
				ActivityID: 1,
				Weekday:    time.Monday,
				Start:      "12:00",
				End:        "13:00",
			},
		},
	}
}
