package mentee

import (
	"alta/immersive-dashboard-api/features/mentees/logs/data"
	"time"
)

type Core struct {
	Id        uint
	CreatedAt time.Time
	UpdateAt time.Time
	DeleteAt time.Time
	FullName        string 
	NickName        string 
	ClassID         uint
	Status          string 
	Category        string 
	Gender          string 
	Graduate        string 
	Mayor           string 
	Phone           string 
	Telegram        string 
	Discord         string 
	Institusi		string
	Email           string 
	EmergencyName   string 
	EmergencyPhone  string 
	EmergencyStatus string 
	Logs						[]data.MenteeLogs `gorm:"foreignKey:MenteeID"`
}