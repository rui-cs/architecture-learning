package domain

import "github.com/rui-cs/architecture-learning/dao"

type UserCompanyOne struct {
	Id int64 `json:"id"`

	Name string `json:"name"`

	Special1 string `json:"special1"`
	Special2 string `json:"special2"`
	Special3 string `json:"special3"`
	Special4 string `json:"special4"`
}

func (u *UserCompanyOne) toDaoUserCompanyOne() *dao.UserCompanyOne {
	return &dao.UserCompanyOne{
		Id:       u.Id,
		Name:     u.Name,
		Special1: u.Special1,
		Special2: u.Special2,
		Special3: u.Special3,
		Special4: u.Special4,
	}
}

func (u *UserCompanyOne) Add(dao dao.CRUDDao) error {
	return dao.Add(u.toDaoUserCompanyOne())
}

func (u *UserCompanyOne) Del(dao dao.CRUDDao) error {
	return dao.Del(u.toDaoUserCompanyOne())
}

func (u *UserCompanyOne) Upd(dao dao.CRUDDao) error {
	return dao.Upd(u.toDaoUserCompanyOne())
}

func (u *UserCompanyOne) Get(dao dao.CRUDDao, cond []interface{}) ([]map[string]interface{}, error) {
	return dao.Get(u.toDaoUserCompanyOne(), cond)
}
