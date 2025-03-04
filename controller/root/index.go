package root

import (
	"net/http"

	"github.com/elvack/billing-engine-api/library/response"
	"github.com/gin-gonic/gin"
)

func (c *controller) Index(ctx *gin.Context) {
	response.Success(ctx, http.StatusOK, "ok", nil)
}
