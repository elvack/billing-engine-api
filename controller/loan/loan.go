package loan

import (
	loanRepo "github.com/elvack/billing-engine-api/repository/loan"
	"github.com/elvack/billing-engine-api/service/loan"
	"gorm.io/gorm"
)

type controller struct {
	loanService loan.IService
}

func NewController(db *gorm.DB) *controller {
	return &controller{loanService: loan.NewService(loanRepo.NewRepo(db))}
}
