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

// GetAllUser implements users.UserServiceInterface
func (service *UserService) GetAllUser() ([]users.Core, error) {
	allUsers, err := service.userData.SelectAll()
	if err != nil {
		return nil, err
	}
	return allUsers, err
}

// DeleteUser implements users.UserServiceInterface
func (service *UserService) DeleteUser(userId uint) error {
	if errDelete := service.userData.Delete(userId); errDelete != nil {
		return errDelete
	}
	return nil
}

// LoginUser implements users.UserServiceInterface
func (service *UserService) LoginUser(email string, password string) (int, error) {
	loginInput := users.LoginUser{
		Email: email,
		Password: password,
	}
	if errValidate := service.validate.Struct(loginInput); errValidate != nil {
		return 0, errValidate
	}
	userId, err := service.userData.Login(loginInput.Email, loginInput.Password)
	if err != nil {
		return 0, err
	}

	return userId, nil
}


func New(userData users.UserDataInterface) users.UserServiceInterface {
	return &UserService{
		userData: userData,
		validate: validator.New(),
	}
}