package admin

import (
	adminRepo "github.com/elvack/billing-engine-api/repository/admin"
	"github.com/elvack/billing-engine-api/service/admin"
	"gorm.io/gorm"
)

type controller struct {
	adminService admin.IService
}

func NewController(db *gorm.DB) *controller {
	return &controller{adminService: admin.NewService(adminRepo.NewRepo(db))}
}
