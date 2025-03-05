package loan

import (
	loanModel "github.com/elvack/billing-engine-api/model/loan"
	"gorm.io/gorm"
)

type (
	IRepo interface {
		Create(loan *loanModel.Loan) (err error)
	}

	repo struct {
		db *gorm.DB
	}
)

func NewRepo(db *gorm.DB) IRepo {
	return &repo{db: db}
}
