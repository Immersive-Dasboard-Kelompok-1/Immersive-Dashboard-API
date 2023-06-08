package data

import (
	"alta/immersive-dashboard-api/features/mentees/mentee"

	"gorm.io/gorm"
)

type MenteeData struct {
	db *gorm.DB
}

// Insert implements mentee.MenteeDataInterface
func (repo *MenteeData) Insert(data mentee.Core) (menteeId uint, err error) {
	menteeData := CoreToMenteeModel(data)
	if tx := repo.db.Create(&menteeData); tx.Error != nil {
		return 0, tx.Error
	} 
	return menteeData.ID, nil
}

// Select implements mentee.MenteeDataInterface
func (repo *MenteeData) Select(menteeId uint) (mentee *mentee.Core, err error) {
	var menteeData Mentees
	if tx := repo.db.First(&menteeData, menteeId); tx.Error != nil {
		return nil, tx.Error
	}
	menteeMap := MenteeModelToCore(menteeData)
	return &menteeMap, nil
}

// SelectAll implements mentee.MenteeDataInterface
func (repo *MenteeData) SelectAll() (mentees []mentee.Core, err error) {
	var menteesData []Mentees
	if tx := repo.db.Find(&menteesData).Where("deleted_at IS NULL"); tx.Error != nil {
		return nil, tx.Error
	}
	var menteesMap []mentee.Core
	for _, mentee := range menteesData {
		menteeCore := MenteeModelToCore(mentee)
		menteesMap = append(menteesMap, menteeCore)
	}
	return menteesMap, nil
}

// Update implements mentee.MenteeDataInterface
func (repo *MenteeData) Update(menteeId uint, data mentee.Core) (mentee *mentee.Core, err error) {
	var menteeData Mentees
	if tx := repo.db.First(&menteeData, menteeId); tx.Error != nil {
		return nil, tx.Error
	}

	menteeMap := CoreToMenteeModel(data)
	if tx := repo.db.Model(&menteeData).Updates(menteeMap); tx.Error != nil {
		return nil, tx.Error
	}
	return &data, nil
}

// Delete implements mentee.MenteeDataInterface
func (repo *MenteeData) Delete(menteeId uint) error {
 if tx := repo.db.Delete(&Mentees{}, menteeId); tx.Error != nil {
	return tx.Error
 }

 return nil
}





func New(db *gorm.DB) mentee.MenteeDataInterface {
	return &MenteeData{
		db: db,
	}
}
