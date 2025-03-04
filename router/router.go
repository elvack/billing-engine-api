package router

import (
	"github.com/elvack/billing-engine-api/controller/admin"
	"github.com/elvack/billing-engine-api/controller/health"
	"github.com/elvack/billing-engine-api/controller/root"
	"github.com/elvack/billing-engine-api/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(db database.DB) (err error) {
	adminController := admin.NewController(db.GormDb)
	healthController := health.NewController(db.SqlDb)
	rootController := root.NewController()
	router := gin.Default()
	router.Use(cors.New(corsConfig))
	router.GET("", rootController.Index)
	adminGroup := router.Group("admin")
	{
		adminGroup.POST("sign-in", adminController.SignIn)
	}
	router.GET("health", healthController.Check)
	router.Static("public", "./public")
	return router.Run()
}
