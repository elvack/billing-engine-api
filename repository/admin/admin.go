package admin

import (
	adminModel "github.com/elvack/billing-engine-api/model/admin"
	"gorm.io/gorm"
)

type (
	IRepo interface {
		Create(admin *adminModel.Admin) (err error)
		Take(selectParams []string, conditions *adminModel.Admin) (admin adminModel.Admin, err error)
	}

	repo struct {
		db *gorm.DB
	}
)

func NewRepo(db *gorm.DB) IRepo {
	return &repo{db: db}
}
