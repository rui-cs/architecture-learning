package dao

import (
	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (u *UserDao) Add(data CRUDDaoDomain) error {
	return data.Add(u.db)
}

func (u *UserDao) Del(data CRUDDaoDomain) error {
	return data.Del(u.db)
}

func (u *UserDao) Upd(data CRUDDaoDomain) error {
	return data.Upd(u.db)
}

func (u *UserDao) Get(data CRUDDaoDomain, cond []interface{}) ([]map[string]interface{}, error) {
	return data.Get(u.db, cond)
}
