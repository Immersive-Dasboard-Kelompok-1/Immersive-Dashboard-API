package data

import (
	"alta/immersive-dashboard-api/features/mentees/mentee/data"

	"gorm.io/gorm"
)

type Classes struct {
	gorm.Model
	Name			string 		`gorm:"type:varchar(50);notNull"`
	Tag				string		`gorm:"type:varchar(5);notNull"`
	UserID		uint
	Mentees 	[]data.Mentees `gorm:"foreignKey:ClassID"`
}