package service

import (
	"errors"
	"fmt"

	errPkg "github.com/pkg/errors"
	"github.com/rui-cs/architecture-learning/dao"
	"github.com/rui-cs/architecture-learning/domain"
)

type UserReq struct {
	Company        int                    `json:"company"`
	ID             int64                  `json:"id"`
	UserCompanyOne *domain.UserCompanyOne `json:"userCompanyOne"`
	UserCompanyTwo *domain.UserCompanyTwo `json:"userCompanyTwo"`
}

type UserService struct {
	UD *dao.UserDao
}

const (
	CompanyOne = 1
	CompanyTwo = 2
)

func NewUserService(UD *dao.UserDao) *UserService {
	return &UserService{UD: UD}
}

var WrongCompany = errors.New("wrong company")

func (r *UserReq) GetCRUDObj() (domain.CRUDObj, error) {
	if r.Company == CompanyOne {
		return r.UserCompanyOne, nil // service 直接操作了domain对象！
	}

	if r.Company == CompanyTwo {
		return r.UserCompanyTwo, nil
	}

	return nil, errPkg.Wrap(WrongCompany, fmt.Sprintf("company : %v", r.Company))
}

func (u *UserService) Add(req *UserReq) error {
	curdObj, err := req.GetCRUDObj()
	if err != nil {
		return err
	}

	return curdObj.Add(u.UD)
}

func (u *UserService) Del(req *UserReq) error {
	curdObj, err := req.GetCRUDObj()
	if err != nil {
		return err
	}

	return curdObj.Del(u.UD)
}

func (u *UserService) Upd(req *UserReq) error {
	curdObj, err := req.GetCRUDObj()
	if err != nil {
		return err
	}

	return curdObj.Upd(u.UD)
}

func (u *UserService) Get(req *UserReq, cond []interface{}) ([]map[string]interface{}, error) {
	curdObj, err := req.GetCRUDObj()
	if err != nil {
		return nil, err
	}

	return curdObj.Get(u.UD, cond)
}
