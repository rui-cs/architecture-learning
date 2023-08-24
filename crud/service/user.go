package service

import (
	"errors"
	"fmt"

	errPkg "github.com/pkg/errors"
	"github.com/rui-cs/architecture-learning/domain"
	"github.com/rui-cs/architecture-learning/repository"
)

type UserReq struct {
	Company        int                    `json:"company"`
	ID             int64                  `json:"id"`
	UserCompanyOne *domain.UserCompanyOne `json:"userCompanyOne"`
	UserCompanyTwo *domain.UserCompanyTwo `json:"userCompanyTwo"`
}

type UserService struct {
	urOne *repository.UserCompanyOneRepo
	urTwo *repository.UserCompanyTwoRepo
}

const (
	CompanyOne = 1
	CompanyTwo = 2
)

func NewUserService(urOne *repository.UserCompanyOneRepo, urTwo *repository.UserCompanyTwoRepo) *UserService {
	return &UserService{urOne: urOne, urTwo: urTwo}
}

var WrongCompany = errors.New("wrong company")

func (r *UserReq) GetCRUDObj() (repository.CRUDObj, error) {
	if r.Company == CompanyOne {
		res := &repository.UserCompanyOneRepo{Data: r.UserCompanyOne}
		res.SetSelf()
		return res, nil
	}

	if r.Company == CompanyTwo {
		res := &repository.UserCompanyTwoRepo{Data: r.UserCompanyTwo}
		res.SetSelf()
		return res, nil
	}

	return nil, errPkg.Wrap(WrongCompany, fmt.Sprintf("company : %v", r.Company))
}

func (u *UserService) Add(req *UserReq) error {
	curdObj, err := req.GetCRUDObj()
	if err != nil {
		return err
	}

	return curdObj.Add()
}

func (u *UserService) Del(req *UserReq) error {
	curdObj, err := req.GetCRUDObj()
	if err != nil {
		return err
	}

	return curdObj.Del()
}

func (u *UserService) Upd(req *UserReq) error {
	curdObj, err := req.GetCRUDObj()
	if err != nil {
		return err
	}

	return curdObj.Upd()
}

func (u *UserService) Get(req *UserReq, cond ...interface{}) ([]map[string]interface{}, error) {
	curdObj, err := req.GetCRUDObj()
	if err != nil {
		return nil, err
	}

	return curdObj.Get(cond)
}
