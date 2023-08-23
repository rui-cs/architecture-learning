package dao

import "gorm.io/gorm"

type UserCompanyTwo struct {
	Id int64 `gorm:"primaryKey,autoIncrement"`

	Name string `gorm:"unique"`

	Special1 string
	Special2 string
}

func (u *UserCompanyTwo) Add(db *gorm.DB) error {
	return db.Create(u).Error
}

func (u *UserCompanyTwo) Del(db *gorm.DB) error {
	return db.Delete(u, u.Id).Error
}

func (u *UserCompanyTwo) Upd(db *gorm.DB) error {
	return db.Model(u).Updates(*u).Error
}

func (u *UserCompanyTwo) Get(db *gorm.DB, cond []interface{}) ([]map[string]interface{}, error) {
	var res []map[string]interface{}
	var err error

	err = db.Model(u).Find(&res, cond...).Error

	return res, err
}
