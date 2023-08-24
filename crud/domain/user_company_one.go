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

func (u *UserCompanyOne) ToDaoUserCompanyOne() *dao.UserCompanyOne {
	return &dao.UserCompanyOne{
		Id:       u.Id,
		Name:     u.Name,
		Special1: u.Special1,
		Special2: u.Special2,
		Special3: u.Special3,
		Special4: u.Special4,
	}
}
