package data

import (
	"alta/immersive-dashboard-api/app/helper"
	"alta/immersive-dashboard-api/features/users"
	"errors"

	"gorm.io/gorm"
)

type UserData struct {
	db *gorm.DB
}

// Insert implements users.UserDataInterface
func (repo *UserData) Insert(data users.Core) (uint, error) {
	hashPassword, err := helper.HashPasword(data.Password)
	if err != nil {
		return 0, errors.New("error hashing password: " + err.Error())
	}

	userData := Users{
		FullName: data.FullName,
		Email:    data.Email,
		Password: hashPassword,
		Team:     data.Team,
		Role:     data.Role,
		Status:   data.Status,
	}

	if tx := repo.db.Create(&userData); tx.Error != nil {
		return 0, tx.Error
	} else if tx.RowsAffected == 0 {
		return 0, errors.New("insert data user failed, rows affected 0 ")
	}
	return userData.ID, nil
}

// Update implements users.UserDataInterface
func (repo *UserData) Update(userId uint, data users.Core) error {
	var user Users
	if tx := repo.db.Where("id = ?", userId).First(&user); tx.Error != nil {
		return tx.Error
	}

	if tx := repo.db.Model(&user).Updates(Users{
		FullName: data.FullName,
		Email:    data.Email,
		Team:     data.Team,
		Role:     data.Role,
		Status:   data.Status,
	}); tx.Error != nil {
		return tx.Error
	}

	return nil
}

// Select implements users.UserDataInterface
func (repo *UserData) Select(userId uint) (users.Core, error) {
	var user Users
	if tx := repo.db.Where("id = ?", userId).First(&user); tx.Error != nil {
		return users.Core{}, tx.Error
	}

	mapUser := users.Core{
		Id:        user.ID,
		FullName:  user.FullName,
		Email:     user.Email,
		Team:      user.Team,
		Role:      user.Role,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return mapUser, nil
}

// SelectAll implements users.UserDataInterface
func (repo *UserData) SelectAll() ([]users.Core, error) {
	var _users []Users
	if tx := repo.db.Find(&_users); tx.Error != nil {
		return nil, tx.Error
	}

	var allUsers []users.Core 
	for _, user := range _users {
		var data = users.Core{
			Id: user.ID,
			FullName: user.FullName,
			Email: user.Email,
			Team: user.Team,
			Role: user.Role,
			Status: user.Status,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		allUsers = append(allUsers, data)
	}

	return allUsers, nil
}

// Delete implements users.UserDataInterface
func (repo *UserData) Delete(userId uint) error {
	if tx := repo.db.Delete(&Users{}, userId); tx.Error != nil {
		return tx.Error
	}

	if errChangeStatusUser := repo.changeStatusUser(userId, "deleted"); errChangeStatusUser != nil {
		return errChangeStatusUser
	}

	return nil
}

// Login implements users.UserDataInterface
func (repo *UserData) Login(email string, password string) (int, error) {
	var user Users
	if tx := repo.db.Where("email = ?", email).First(&user); tx.Error != nil {
		return 0, tx.Error
	}
	
	match := helper.CheckPaswordHash(password, user.Password)
	if !match {
		return 0, errors.New("kredensial tidak cocok")
	}

	if errChangeStatusUser := repo.changeStatusUser(user.ID, "active"); errChangeStatusUser != nil {
		return 0, errChangeStatusUser
	}

	return int(user.ID), nil
}

func (repo *UserData) changeStatusUser(userId uint, status string) error {
	var user Users
	if tx := repo.db.First(&user, userId); tx.Error != nil {
		return tx.Error
	}

	if tx := repo.db.Model(&user).Update("status", status); tx.Error != nil {
		return tx.Error
	}

	return nil
}

func New(db *gorm.DB) users.UserDataInterface {
	return &UserData{
		db: db,
	}
}
