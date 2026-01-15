package models

type User struct {
	UserID          uint          `gorm:"primaryKey" json:"user_id"`
	FirstName       string        `json:"first_name"`
	LastName        string        `json:"last_name"`
	Email           string        `json:"email"`
	Phone           string        `json:"phone"`
	GroupCalendarID uint          `json:"group_calendar_id"`
	GroupCalendar   GroupCalendar `gorm:"foreignKey:GroupCalendarID" json:"-"`
}
