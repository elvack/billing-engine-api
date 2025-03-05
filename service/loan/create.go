package loan

import loanModel "github.com/elvack/billing-engine-api/model/loan"

func (s *service) Create(reqBody *loanModel.CreateReqBody) (err error) {
	loan := loanModel.Loan{
		Amount: 5000000,
		CustomerId: reqBody.CustomerId,
		FlatInterestRate: 10,
		Installment: 50,
		Status: "OPEN",
	}
	return s.loanRepo.Create(&loan)
}
