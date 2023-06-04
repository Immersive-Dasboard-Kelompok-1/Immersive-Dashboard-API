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
func (repo *UserData) Insert(data users.Core) error {
	hashPassword, err := helper.HashPasword(data.Password)
	if err != nil {
		return errors.New("error hashing password: " + err.Error())
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
		return tx.Error
	} else if tx.RowsAffected == 0 {
		return errors.New("insert data user failed, rows affected 0 ")
	}

	return nil
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
		Id: user.ID,
		FullName: user.FullName,
		Email: user.Email,
		Team: user.Team,
		Role: user.Role,
		Status: user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return mapUser, nil
}

// Delete implements users.UserDataInterface
func (*UserData) Delete(userId uint) error {
	panic("unimplemented")
}

// SelectAll implements users.UserDataInterface
func (*UserData) SelectAll() ([]users.Core, error) {
	panic("unimplemented")
}

func New(db *gorm.DB) users.UserDataInterface {
	return &UserData{
		db: db,
	}
}
