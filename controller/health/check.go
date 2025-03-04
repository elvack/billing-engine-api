package health

import (
	"net/http"

	"github.com/elvack/billing-engine-api/common"
	"github.com/elvack/billing-engine-api/library/response"
	"github.com/gin-gonic/gin"
)

func (c *controller) Check(ctx *gin.Context) {
	if err := c.healthService.Check(); err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, common.CheckedSuccessfully, nil)
}
