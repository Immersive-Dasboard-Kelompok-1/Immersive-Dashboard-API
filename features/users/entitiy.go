package users

import "time"

type Core struct {
	Id 					uint 		
	FullName		string 
	Email				string `validate:"required,email"`
	Password		string `validate:"required"`
	Team				string
	Role				string 
	Status			string
	CreatedAt		time.Time
	UpdatedAt		time.Time
}

type LoginUser struct {
	Email			string `validate:"required,email"`
	Password	string `validate:"required"` 
}

type UserDataInterface interface {
	Insert(data Core) error
	Update(userId uint, data Core) error
	Select(userId uint) (Core, error)
	SelectAll() ([]Core, error)
	Delete(userId uint) error
	Login(email string, password string) (int, error)
}

type UserServiceInterface interface {
	AddUser(data Core) error
	EditUser(userId uint, data Core) error
	GetUser(userId uint) (Core, error)
	GetAllUser() ([]Core, error)
	DeleteUser(userId uint) error
	LoginUser(email string, password string) (int, error)
}