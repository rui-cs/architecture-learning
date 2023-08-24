package domain

import "github.com/rui-cs/architecture-learning/dao"

type UserCompanyTwo struct {
	Id int64 `json:"id"`

	Name string `json:"name"`

	Special1 string `json:"special1"`
	Special2 string `json:"special2"`
}

func (u *UserCompanyTwo) ToDaoUserCompanyTwo() *dao.UserCompanyTwo {
	return &dao.UserCompanyTwo{
		Id:       u.Id,
		Name:     u.Name,
		Special1: u.Special1,
		Special2: u.Special2,
	}
}
