package routers

import (
	"api-goklaster/middleware"

	"github.com/gin-gonic/gin"
)

func DashboardRouter(r *gin.Engine) {
	app := r.Group("/v1")
	app.Use(middleware.TokenValidator())
	app.GET("/dashboard", func(c *gin.Context) {
		c.String(200, "dashboard")
	})
}
