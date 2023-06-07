package handler

import "alta/immersive-dashboard-api/features/mentees/logs"

type LogsRequest struct {
	Status   string `json:"status" form:"status"`	
	Feedback string `json:"feedback" form:"feedback"`
	MenteeID uint `json:"mentee_Id" form:"mentee_Id"`
	UserID   uint `json:"user_id" form:"user_Id"`
}

func RequestToCoreLogs(input LogsRequest) logs.Core{
	return logs.Core{
		Status: input.Status,
		Feedback: input.Feedback,
		MenteeID: input.MenteeID,
		UserID: input.UserID,
	}
}
