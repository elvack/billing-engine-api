package loan

import loanModel "github.com/elvack/billing-engine-api/model/loan"

func (r *repo) Create(loan *loanModel.Loan) (err error) {
	return r.db.Create(loan).Error
}
