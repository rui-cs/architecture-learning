package dao

import "gorm.io/gorm"

type UserCompanyOne struct {
	Id int64 `gorm:"primaryKey,autoIncrement"`

	Name string `gorm:"unique"`

	Special1 string
	Special2 string
	Special3 string
	Special4 string
}

type UserCompanyOneDao struct {
	db *gorm.DB
}

func NewUserCompanyOneDao(db *gorm.DB) *UserCompanyOneDao {
	return &UserCompanyOneDao{db: db}
}

func (u *UserCompanyOneDao) Add(data *UserCompanyOne) error {
	return u.db.Create(data).Error
}

func (u *UserCompanyOneDao) Del(data *UserCompanyOne) error {
	return u.db.Delete(data, data.Id).Error
}

func (u *UserCompanyOneDao) Upd(data *UserCompanyOne) error {
	return u.db.Model(data).Updates(data).Error
}

func (u *UserCompanyOneDao) Get(cond []interface{}) ([]map[string]interface{}, error) {
	var res []map[string]interface{}
	err := u.db.Model(&UserCompanyOne{}).Find(&res, cond).Error

	return res, err
}
