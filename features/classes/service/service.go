package service

import (
	//"be17/main/feature/user"

	"alta/immersive-dashboard-api/features/classes"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type classService struct {
	classData classes.ClassDataInterface
	validate  *validator.Validate
}

// Edit implements classes.ClassServiceInterface
func (service *classService) Edit(id int, input classes.Core) error {
	err := service.classData.Update(id,input)
	if err != nil{
		return fmt.Errorf("failed to update classses with ID %d:%w",id,err)
	}
	return nil
}

// Create implements classes.ClassServiceInterface
func (service *classService) Create(input classes.Core) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}

	// Menyimpan data user ke database
	errInsert := service.classData.Insert(input)
	return errInsert
}

func New(repo classes.ClassDataInterface) classes.ClassServiceInterface {
	return &classService{
		classData: repo,
		validate:  validator.New(),
	}
}