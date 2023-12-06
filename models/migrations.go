package models

import "gorm.io/gorm"

func Migrate(db *gorm.DB) (err error) {
	err = db.AutoMigrate(&WelderTest{}, &MarkerTest{})
	if err != nil {
		return
	}
	return
}
