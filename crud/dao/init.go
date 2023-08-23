package dao

import "gorm.io/gorm"

func InitTable(db *gorm.DB) error {
	if err := db.AutoMigrate(&UserCompanyOne{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&UserCompanyTwo{}); err != nil {
		return err
	}

	return nil
}
