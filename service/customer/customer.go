package customer

import (
	customerModel "github.com/elvack/billing-engine-api/model/customer"
	customerRepo "github.com/elvack/billing-engine-api/repository/customer"
)

type (
	IService interface {
		Create(reqBody *customerModel.CreateReqBody) (statusCode int, err error)
	}

	service struct {
		customerRepo customerRepo.IRepo
	}
)

func NewService(customerRepo customerRepo.IRepo) IService {
	return &service{customerRepo: customerRepo}
}
