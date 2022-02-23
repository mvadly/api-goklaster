package v1

import (
	"api-goklaster/v1/routers"

	"github.com/gin-gonic/gin"
)

func GinRoutes() *gin.Engine {
	router := gin.Default()
	routers.AuthRouter(router)
	routers.DashboardRouter(router)
	return router
}
