package health

import healthRepo "github.com/elvack/billing-engine-api/repository/health"

type (
	IService interface {
		Check() (err error)
	}

	service struct {
		healthRepo healthRepo.IRepo
	}
)

func NewService(healthRepo healthRepo.IRepo) IService {
	return &service{healthRepo: healthRepo}
}
