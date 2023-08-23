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

func (u *UserCompanyOne) Add(db *gorm.DB) error {
	return db.Create(u).Error
}

func (u *UserCompanyOne) Del(db *gorm.DB) error {
	return db.Delete(u, u.Id).Error
}

func (u *UserCompanyOne) Upd(db *gorm.DB) error {
	return db.Model(u).Updates(*u).Error
}

func (u *UserCompanyOne) Get(db *gorm.DB, cond []interface{}) ([]map[string]interface{}, error) {
	var res []map[string]interface{}
	err := db.Model(u).Find(&res, cond).Error

	return res, err
}

//func (u *UserCompanyOne) ToMap() map[string]interface{} {
//	return map[string]interface{}{
//		"id":   u.Id,
//		"name": u.Name,
//	}
//}
