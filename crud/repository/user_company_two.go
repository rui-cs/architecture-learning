package repository

import (
	"github.com/rui-cs/architecture-learning/dao"
	"github.com/rui-cs/architecture-learning/domain"
)

type UserCompanyTwoRepo struct {
	Data *domain.UserCompanyTwo
	ud   *dao.UserCompanyTwoDao
}

func NewUserCompanyTwoRepo(data *domain.UserCompanyTwo, ud *dao.UserCompanyTwoDao) *UserCompanyTwoRepo {
	return &UserCompanyTwoRepo{Data: data, ud: ud}
}

var gUserCompanyTwoRepo *UserCompanyTwoRepo

func (u *UserCompanyTwoRepo) SetSelf() {
	u.ud = gUserCompanyTwoRepo.ud
}

func (u *UserCompanyTwoRepo) Add() error {
	return u.ud.Add(u.Data.ToDaoUserCompanyTwo())
}

func (u *UserCompanyTwoRepo) Del() error {
	return u.ud.Del(u.Data.ToDaoUserCompanyTwo())
}

func (u *UserCompanyTwoRepo) Upd() error {
	return u.ud.Upd(u.Data.ToDaoUserCompanyTwo())
}

func (u *UserCompanyTwoRepo) Get(cond []interface{}) ([]map[string]interface{}, error) {
	return u.ud.Get(cond)
}
