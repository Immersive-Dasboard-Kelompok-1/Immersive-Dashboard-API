package data

import (
	"alta/immersive-dashboard-api/features/classes"
	"errors"

	"gorm.io/gorm"
)

type classQuery struct {
	db *gorm.DB
}

// Delete implements classes.ClassDataInterface
func (repo *classQuery) Deleted(id int,UserId int) error {
	var classData Classes
	errDelete := repo.db.Delete(&classData,"id=? AND user_id=?",id,UserId)
	if errDelete.Error != nil{
		return errDelete.Error
	}
	return nil
}

// Update implements classes.ClassDataInterface
func (repo *classQuery) Update(id int,UserId int, input classes.Core) error {
	classInput := CoreToModel(input)
	err := repo.db.Model(&Classes{}).Where("id=? AND user_id=?",id,UserId).Updates(UpdateClass(classInput))
	if err != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("no rows affected, update failed")
	}
	return nil
}

// Insert implements classes.ClassDataInterface
func (repo *classQuery) Insert(input classes.Core,UserId int) error {
	classInput := CoreToModel(input)
	classInput.UserID = uint(UserId)
	tx := repo.db.Create(&classInput)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failes, row affected = 0")
	}
	return nil
}

func New(db *gorm.DB) classes.ClassDataInterface {
	return &classQuery{
		db: db,
	}
}
