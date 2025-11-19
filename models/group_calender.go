package models

type GroupCalender struct {
	GroupCalenderID uint   `gorm:"primaryKey" json:"grup_calender_id"`
	Name            string `json:"name"`
	WorkStart       string `json:"work_start"`
	WorkEnd         string `json:"work_end"`
	WeekendBlocked  bool   `json:"weekend_blocked"`
	Members         []User `gorm:"foreignKey:GroupCalenderID" json:"members,omitempty"`
}
