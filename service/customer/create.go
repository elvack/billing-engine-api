package customer

import (
	"errors"
	"net/http"

	"github.com/elvack/billing-engine-api/common"
	customerModel "github.com/elvack/billing-engine-api/model/customer"
	"gorm.io/gorm"
)

func (s *service) Create(reqBody *customerModel.CreateReqBody) (statusCode int, err error) {
	if _, err = s.customerRepo.Take([]string{"id"}, &customerModel.Customer{PhoneNumber: reqBody.PhoneNumber}); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusInternalServerError, err
		}
	} else {
		return http.StatusBadRequest, errors.New(common.PhoneNumberAlreadyExists)
	}
	customer := customerModel.Customer{
		Name: reqBody.Name,
		PhoneNumber: reqBody.PhoneNumber,
	}
	return http.StatusOK, s.customerRepo.Create(&customer)
}
