package customer

type CreateReqBody struct {
	Name	    string `binding:"required"`
	PhoneNumber string `binding:"gte=10" json:"phone_number"`
}
