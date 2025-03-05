package loan

import "time"

type Loan struct {
	Amount           int64
	CreatedAt        time.Time
	CustomerId       uint32
	FlatInterestRate int16
	Id               uint32
	Installment      int16
	Outstanding      int64
	Status           string
	UpdatedAt        time.Time
}
