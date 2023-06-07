package service

import (
	"alta/immersive-dashboard-api/features/mentees/logs"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type LogsService struct {
	logsData logs.LogsDataInterface
	validate *validator.Validate
}

// Edit implements logs.LogsServiceInterface
func (service *LogsService) Edit(input logs.Core, id uint) error {
	err := service.logsData.Update(input,id)
	if err != nil {
		return fmt.Errorf("failed to update classses with ID %d:%w", id, err)
	}
	return nil
}

// Add implements logs.LogsServiceInterface
func (service *LogsService) Add(input logs.Core, userId uint) error {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}
	err := service.logsData.Insert(input, userId)
	if err != nil {
		return err
	}
	return nil
}

func New(logsData logs.LogsDataInterface) logs.LogsServiceInterface {
	return &LogsService{
		logsData: logsData,
		validate: validator.New(),
	}
}
