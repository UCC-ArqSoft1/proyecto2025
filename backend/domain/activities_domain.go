package domain

import "time"

// Activity representa una actividad periódica del gimnasio (por ejemplo Zumba).
type Activity struct {
	ID          int        `json:"id" gorm:"primary_key"`
	Name        string     `json:"name"`                                                              // Ej: "Zumba", "Musculación"
	Description string     `json:"description" gorm:"not_null"`                                       // Opcional
	Schedule    []TimeSlot `json:"schedule" gorm:"foreignKey:ActivityID;constraint:OnDelete:CASCADE"` // Días y horarios en que se repite la actividad
}

// TimeSlot representa un horario recurrente en un día de la semana.
type TimeSlot struct {
	ID         int          `json:"id" gorm:"primaryKey"`
	ActivityID int          `json:"activity_id" gorm:"not null;index"` // Foreign key
	Weekday    time.Weekday `json:"weekday" gorm:"not null"`           // Ej: time.Monday
	Start      string       `json:"start" gorm:"not null"`             // Ej: "18:00"
	End        string       `json:"end" gorm:"not null"`               // Ej: "19:00"
}
