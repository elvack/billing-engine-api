package loan

import (
	"net/http"

	"github.com/elvack/billing-engine-api/common"
	"github.com/elvack/billing-engine-api/library/response"
	loanModel "github.com/elvack/billing-engine-api/model/loan"
	"github.com/gin-gonic/gin"
)

func (c *controller) Create(ctx *gin.Context) {
	var reqBody loanModel.CreateReqBody
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if err := c.loanService.Create(&reqBody); err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, common.CreatedSuccessfully, nil)
}
