package admin

import adminModel "github.com/elvack/billing-engine-api/model/admin"

func (r *repo) Take(selectParams []string, conditions *adminModel.Admin) (admin adminModel.Admin, err error) {
	return admin, r.db.Select(selectParams).Take(&admin, conditions).Error
}
