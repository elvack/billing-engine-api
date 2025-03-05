package loan

type CreateReqBody struct {
	CustomerId uint32 `binding:"required" json:"customer_id"`
}
