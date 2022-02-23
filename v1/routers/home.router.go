package routers

import (
	"api-goklaster/v1/controllers"

	"github.com/gin-gonic/gin"
)

func HomeRouter(r *gin.Engine) {
	app := r.Group("/v1")
	{
		app.GET("/home/get_banners", controllers.GetBanners)
	}

}
