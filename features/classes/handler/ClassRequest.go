package handler

import (
	"alta/immersive-dashboard-api/features/classes"
)

type ClassRequest struct {
	Name   string `json:"name" form:"name"`
	Tag    string `json:"tag" form:"tag"`
	UserID int    `json:"user_id" form:"user_id"`
}

func RequestToCore(input ClassRequest) classes.Core{
	return classes.Core{
		Name: input.Name,
		Tag: input.Tag,
		UserID: uint(input.UserID),
	}
}