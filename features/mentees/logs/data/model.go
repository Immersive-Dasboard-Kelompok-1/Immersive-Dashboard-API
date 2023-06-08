package data

import (
	"alta/immersive-dashboard-api/features/mentees/logs"

	"gorm.io/gorm"
)

type MenteeLogs struct {
	gorm.Model
	Status 			string 	`gorm:"type:varchar(50)"`
	Feedback		string	`gorm:"type:text"`
	MenteeID		uint
	UserID			uint
}

func CoreToModelLogs(input logs.Core) MenteeLogs{
	return MenteeLogs{
		Status: input.Status,
		Feedback: input.Feedback,
		MenteeID: input.MenteeID,
		UserID: input.UserID,
	}
}

func LogsModelToCore(input MenteeLogs) logs.Core{
	return logs.Core{
		Id: input.ID,
		Status:  input.Status,
		Feedback: input.Feedback,
		MenteeID: input.MenteeID,
		UserID:   input.UserID,
	}
}