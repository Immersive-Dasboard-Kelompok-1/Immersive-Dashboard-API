package data

import (
	"alta/immersive-dashboard-api/features/mentees/logs/data"

	"gorm.io/gorm"
)

type Mentees struct {
	gorm.Model
	FullName 				string 	`gorm:"type:varchar(100)"`
	NickName				string	`gorm:"type:varchar(10)"`
	ClassID					uint
	Status					string	`gorm:"type:varchar(50)"`
	Category				string 	`gorm:"type:enum('it','non-it')"`
	Gender					string	`gorm:"type:enum('male','female')"`
	Graduate				string	`gorm:"type:varchar(50)"`
	Mayor						string	`gorm:"type:varchar(50)"`
	Phone						string 	`gorm:"type:varchar(50)"`
	Telegram				string	`gorm:"type:varchar(50)"`
	Discord					string	`gorm:"type:varchar(50)"`
	Email						string	`gorm:"type:varchar(50)"`
	EmergencyName		string	`gorm:"type:varchar(50)"`
	EmergencyPhone	string	`gorm:"type:varchar(50)"`
	EmergencyStatus	string	`gorm:"type:varchar(50)"`
	Logs						[]data.MenteeLogs `gorm:"foreignKey:MenteeID"`
}