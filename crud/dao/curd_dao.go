package dao

import "gorm.io/gorm"

type CRUDDao interface {
	Add(data CRUDDaoDomain) error
	Del(data CRUDDaoDomain) error
	Upd(data CRUDDaoDomain) error
	Get(data CRUDDaoDomain, cond []interface{}) ([]map[string]interface{}, error)
}

// 遇到if-else考虑用interface{}抽象
type CRUDDaoDomain interface {
	Add(db *gorm.DB) error
	Del(db *gorm.DB) error
	Upd(db *gorm.DB) error
	Get(db *gorm.DB, cond []interface{}) ([]map[string]interface{}, error)
}
