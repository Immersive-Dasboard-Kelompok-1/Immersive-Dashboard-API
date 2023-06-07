package handler

import "alta/immersive-dashboard-api/features/mentees/logs"

type LogsRequest struct {
	Status   string `json:"proof" form:"proof"`	
	Feedback string `json:"notes" form:"notes"`
	MenteeID uint `json:"id_mentee" form:"id_mentee"`
	UserID   uint `json:"id_user" form:"id_user"`
}

func RequestToCoreLogs(input LogsRequest) logs.Core{
	return logs.Core{
		Status: input.Status,
		Feedback: input.Feedback,
		MenteeID: input.MenteeID,
		UserID: input.UserID,
	}
}
