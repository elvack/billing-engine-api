package customer

import (
	customerRepo "github.com/elvack/billing-engine-api/repository/customer"
	"github.com/elvack/billing-engine-api/service/customer"
	"gorm.io/gorm"
)

type controller struct {
	customerService customer.IService
}

func NewController(db *gorm.DB) *controller {
	return &controller{customerService: customer.NewService(customerRepo.NewRepo(db))}
}
