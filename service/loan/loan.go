package loan

import (
	loanModel "github.com/elvack/billing-engine-api/model/loan"
	loanRepo "github.com/elvack/billing-engine-api/repository/loan"
)

type (
	IService interface {
		Create(reqBody *loanModel.CreateReqBody) (err error)
	}

	service struct {
		loanRepo loanRepo.IRepo
	}
)

func NewService(loanRepo loanRepo.IRepo) IService {
	return &service{loanRepo: loanRepo}
}
