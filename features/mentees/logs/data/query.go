package data

import (
	"alta/immersive-dashboard-api/features/mentees/logs"
	"errors"

	"gorm.io/gorm"
)

type LogsData struct {
	db *gorm.DB
}

// Update implements logs.LogsDataInterface
func (repo *LogsData) Update(input logs.Core, id uint) error {
	var logs MenteeLogs
	
	if tx := repo.db.Model(&logs).Where("id=?",id).Updates(CoreToModelLogs(input)); tx.Error != nil {
		return tx.Error
	}

	return nil
}

// Insert implements logs.LogsDataInterface
func (repo *LogsData) Insert(input logs.Core, userId uint) error {
	logsInput := CoreToModelLogs(input)
	logsInput.UserID = userId
	tx := repo.db.Create(&logsInput)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failes, row affected = 0")
	}
	return nil
}

func New(db *gorm.DB) logs.LogsDataInterface {
	return &LogsData{
		db: db,
	}
}
