package customer

import customerModel "github.com/elvack/billing-engine-api/model/customer"

func (r *repo) Create(customer *customerModel.Customer) (err error) {
	return r.db.Create(customer).Error
}
