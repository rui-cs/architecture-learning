package dao

import "gorm.io/gorm"

type UserCompanyTwo struct {
	Id int64 `gorm:"primaryKey,autoIncrement"`

	Name string `gorm:"unique"`

	Special1 string
	Special2 string
}

type UserCompanyTwoDao struct {
	db *gorm.DB
}

func NewUserCompanyTwoDao(db *gorm.DB) *UserCompanyTwoDao {
	return &UserCompanyTwoDao{db: db}
}

func (u *UserCompanyTwoDao) Add(data *UserCompanyTwo) error {
	return u.db.Create(data).Error
}

func (u *UserCompanyTwoDao) Del(data *UserCompanyTwo) error {
	return u.db.Delete(data, data.Id).Error
}

func (u *UserCompanyTwoDao) Upd(data *UserCompanyTwo) error {
	return u.db.Model(data).Updates(data).Error
}

func (u *UserCompanyTwoDao) Get(cond []interface{}) ([]map[string]interface{}, error) {
	var res []map[string]interface{}
	err := u.db.Model(&UserCompanyTwo{}).Find(&res, cond).Error

	return res, err
}
