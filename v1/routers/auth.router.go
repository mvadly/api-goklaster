package routers

import (
	"api-goklaster/v1/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.Engine) {
	app := r.Group("/v1")
	{
		app.POST("/login", controllers.Login)
	}

}
