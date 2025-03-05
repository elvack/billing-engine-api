package customer

import (
	"net/http"

	"github.com/elvack/billing-engine-api/common"
	"github.com/elvack/billing-engine-api/library/response"
	customerModel "github.com/elvack/billing-engine-api/model/customer"
	"github.com/gin-gonic/gin"
)

func (c *controller) Create(ctx *gin.Context) {
	var reqBody customerModel.CreateReqBody
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	statusCode, err := c.customerService.Create(&reqBody)
	if err != nil {
		response.Error(ctx, statusCode, err.Error())
		return
	}
	response.Success(ctx, statusCode, common.CreatedSuccessfully, nil)
}
