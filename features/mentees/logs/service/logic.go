package service

import (
	"alta/immersive-dashboard-api/features/mentees/logs"

	"github.com/go-playground/validator/v10"
)

type ServiceData struct {
	LogsData logs.LogsDataInterface
	validate *validator.Validate
}

func New( Logs)