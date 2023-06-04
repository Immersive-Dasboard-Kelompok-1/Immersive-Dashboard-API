package data

import "gorm.io/gorm"

type MenteeLogs struct {
	gorm.Model
	Status 			string 	`gorm:"type:varchar(50)"`
	Feedback		string	`gorm:"type:text"`
	MenteeID		uint
	UserID			uint
}