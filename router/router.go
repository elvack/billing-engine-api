package router

import (
	"github.com/elvack/billing-engine-api/controller/health"
	"github.com/elvack/billing-engine-api/controller/root"
	"github.com/elvack/billing-engine-api/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(db database.DB) (err error) {
	healthController := health.NewController(db.SqlDb)
	rootController := root.NewController()
	router := gin.Default()
	router.Use(cors.New(corsConfig))
	router.GET("", rootController.Index)
	router.GET("health", healthController.Check)
	router.Static("public", "./public")
	return router.Run()
}
