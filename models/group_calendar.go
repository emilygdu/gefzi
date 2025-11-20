package models

type GroupCalendar struct {
	GroupCalendarID uint   `gorm:"primaryKey" json:"grup_calendar_id"`
	Name            string `json:"name"`
	WorkStart       string `json:"work_start"`
	WorkEnd         string `json:"work_end"`
	WeekendBlocked  bool   `json:"weekend_blocked"`
	Members         []User `gorm:"foreignKey:GroupCalendarID" json:"members,omitempty"`
}

// Ausblick man könnte wenn Kalender erstellbar sind noch eine Auswahl von den Symbolen hinzufügen
