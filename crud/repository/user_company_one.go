package repository

import (
	"github.com/rui-cs/architecture-learning/dao"
	"github.com/rui-cs/architecture-learning/domain"
)

type UserCompanyOneRepo struct {
	Data *domain.UserCompanyOne
	ud   *dao.UserCompanyOneDao
}

func Init(udOne *dao.UserCompanyOneDao, udTwo *dao.UserCompanyTwoDao) {
	gUserCompanyOneRepo = &UserCompanyOneRepo{
		ud: udOne,
	}

	gUserCompanyTwoRepo = &UserCompanyTwoRepo{
		ud: udTwo,
	}
}

var gUserCompanyOneRepo *UserCompanyOneRepo

func NewUserCompanyOneRepo(data *domain.UserCompanyOne, ud *dao.UserCompanyOneDao) *UserCompanyOneRepo {
	return &UserCompanyOneRepo{Data: data, ud: ud}
}

func (u *UserCompanyOneRepo) SetSelf() {
	u.ud = gUserCompanyOneRepo.ud
}

func (u *UserCompanyOneRepo) Add() error {
	return u.ud.Add(u.Data.ToDaoUserCompanyOne())
}

func (u *UserCompanyOneRepo) Del() error {
	return u.ud.Del(u.Data.ToDaoUserCompanyOne())
}

func (u *UserCompanyOneRepo) Upd() error {
	return u.ud.Upd(u.Data.ToDaoUserCompanyOne())
}

func (u *UserCompanyOneRepo) Get(cond []interface{}) ([]map[string]interface{}, error) {
	return u.ud.Get(cond)
}
