package domain

import "github.com/rui-cs/architecture-learning/dao"

type UserCompanyTwo struct {
	Id int64 `json:"id"`

	Name string `json:"name"`

	Special1 string `json:"special1"`
	Special2 string `json:"special2"`
}

func (u *UserCompanyTwo) toDaoUserCompanyTwo() *dao.UserCompanyTwo {
	return &dao.UserCompanyTwo{
		Id:       u.Id,
		Name:     u.Name,
		Special1: u.Special1,
		Special2: u.Special2,
	}
}

func (u *UserCompanyTwo) Add(dao dao.CRUDDao) error {
	return dao.Add(u.toDaoUserCompanyTwo())
}

func (u *UserCompanyTwo) Del(dao dao.CRUDDao) error {
	return dao.Del(u.toDaoUserCompanyTwo())
}

func (u *UserCompanyTwo) Upd(dao dao.CRUDDao) error {
	return dao.Upd(u.toDaoUserCompanyTwo())
}

func (u *UserCompanyTwo) Get(dao dao.CRUDDao, cond []interface{}) ([]map[string]interface{}, error) {
	return dao.Get(u.toDaoUserCompanyTwo(), cond)
}
