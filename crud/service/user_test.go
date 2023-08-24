package service

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/rui-cs/architecture-learning/domain"
)

func TestPrintReqJson(t *testing.T) {
	req := UserReq{
		Company: CompanyOne,
		ID:      1,
		UserCompanyOne: &domain.UserCompanyOne{
			Id:       1,
			Name:     "user1",
			Special1: "special1",
			Special2: "special2",
			Special3: "special3",
			Special4: "special4",
		},
		UserCompanyTwo: &domain.UserCompanyTwo{
			Id:       1,
			Name:     "user2",
			Special1: "special1",
			Special2: "special2",
		},
	}

	b, _ := json.Marshal(req)
	fmt.Print(string(b))
}
