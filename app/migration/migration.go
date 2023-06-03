package migration

import "gorm.io/gorm"

func InitMigrate(db *gorm.DB) {
	// TODO: Tambahkan setiap model fitur didalam parameter pisahkan dengan koma
	db.AutoMigrate()
}