package models

import (
	"time"
)

type Event struct {
	ID          uint      `gorm:"primaryKey"`
	Date        string    `json:"date"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Visibillity string    `json:"visibility"` //private oder business
	//User-Termin
	UserID uint `json:"user_id"`
	//Gruppen-Termin
	GroupCalendarID uint `json:"group_calendar_id"`
}

//Begründung warum wir keine Termin Beschreibung haben, Begründung auch für Termin Titel
