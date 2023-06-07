package data

import (
	classData "alta/immersive-dashboard-api/features/classes/data"
	logData "alta/immersive-dashboard-api/features/mentees/logs/data"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	FullName		string					`gorm:"type:varchar(100)"` 
	Email				string					`gorm:"type:varchar(50);unique;notNull" validate:"required,email"`
	Password		string					`gorm:"type:varchar(500);notNull"`
	Team 				string 					`gorm:"type:varchar(50)"`
	Role				string 					`gorm:"type:enum('user','admin');default:'user'"`
	Status			string 					`gorm:"type:enum('active','not-active','deleted');default:'active'"`
	Classes			[]classData.Classes	`gorm:"foreignKey:UserID"`
	Logs				[]logData.MenteeLogs `gorm:"foreignKey:UserID"`
}

