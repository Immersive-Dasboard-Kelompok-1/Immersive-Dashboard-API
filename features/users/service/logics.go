package service

import (
	"alta/immersive-dashboard-api/features/users"

	"github.com/go-playground/validator/v10"
)

type UserService struct {
	userData users.UserDataInterface
	validate *validator.Validate
}

// AddUser implements users.UserServiceInterface
func (service *UserService) AddUser(data users.Core) error {
	if errValidate := service.validate.Struct(data); errValidate != nil {
		return errValidate
	}
	if err := service.userData.Insert(data); err != nil {
		return err
	}
	return nil
}

// EditUser implements users.UserServiceInterface
func (service *UserService) EditUser(userId uint, data users.Core) error {
	if errValidate := service.validate.Struct(data); errValidate != nil {
		return errValidate
	}
	if err := service.userData.Update(userId, data); err != nil {
		return err
	}
	return nil
}

// GetUser implements users.UserServiceInterface
func (repo *UserService) GetUser(userId uint) (users.Core, error) {
	user, err := repo.userData.Select(userId)
	if err != nil {
		return users.Core{}, err
	}
	return user, err
}

// DeleteUser implements users.UserServiceInterface
func (*UserService) DeleteUser() error {
	panic("unimplemented")
}

// GetAllUser implements users.UserServiceInterface
func (*UserService) GetAllUser() ([]users.Core, error) {
	panic("unimplemented")
}

func New(userData users.UserDataInterface) users.UserServiceInterface {
	return &UserService{
		userData: userData,
		validate: validator.New(),
	}
}
