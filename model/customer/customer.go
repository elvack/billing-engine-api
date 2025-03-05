package customer

import "time"

type Customer struct {
	CreatedAt    time.Time
	Id           uint32
	IsBorrower   bool
	IsDelinquent bool
	Name         string
	PhoneNumber  string
	UpdatedAt    time.Time
}
