package customer

import (
	customerModel "github.com/elvack/billing-engine-api/model/customer"
	"gorm.io/gorm"
)

type (
	IRepo interface {
		Create(customer *customerModel.Customer) (err error)
		Take(selectParams []string, conditions *customerModel.Customer) (customer customerModel.Customer, err error)
	}

	repo struct {
		db *gorm.DB
	}
)

func NewRepo(db *gorm.DB) IRepo {
	return &repo{db: db}
}
