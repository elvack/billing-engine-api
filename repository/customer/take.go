package customer

import customerModel "github.com/elvack/billing-engine-api/model/customer"

func (r *repo) Take(selectParams []string, conditions *customerModel.Customer) (customer customerModel.Customer, err error) {
	return customer, r.db.Select(selectParams).Take(&customer, conditions).Error
}
