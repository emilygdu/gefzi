package models

type Event struct {
	EventID     uint   `gorm:"primaryKey" json:"event_id"`
	Date        string `json:"date"`       // YYYY-MM-DD
	StartTime   string `json:"start_time"` // HH:MM
	EndTime     string `json:"end_time"`   // HH:MM
	Visibillity string `json:"visibility"` //private oder business
	//User-Termin
	UserID uint `json:"user_id"`
	//Gruppen-Termin
	GroupCalendarID uint `json:"group_calendar_id"`
}
