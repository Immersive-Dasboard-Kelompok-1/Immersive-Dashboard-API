package service

import (
	"alta/immersive-dashboard-api/features/mentees/logs"

	"github.com/go-playground/validator/v10"
)

type LogsService struct {
	logsData logs.LogsDataInterface
	validate *validator.Validate
}

// Add implements logs.LogsServiceInterface
func (service *LogsService) Add(input logs.Core, userId uint) error {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}
	err := service.logsData.Insert(input,userId); if err != nil {
		return  err
	}
	return  nil
}

func New(logsData logs.LogsDataInterface) logs.LogsServiceInterface {
	return &LogsService{
		logsData: logsData,
		validate: validator.New(),
	}
}
