package domain

import "github.com/rui-cs/architecture-learning/dao"

type CRUDObj interface {
	Add(dao dao.CRUDDao) error
	Del(dao dao.CRUDDao) error
	Upd(dao dao.CRUDDao) error
	Get(dao dao.CRUDDao, cond []interface{}) ([]map[string]interface{}, error)
}
