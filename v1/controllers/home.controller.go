package controllers

import (
	"api-goklaster/v1/repositories"

	"github.com/gin-gonic/gin"
)

func GetBanners(c *gin.Context) {
	// c.String(200, "banner")
	var banners, errBanner = repositories.GetBanners()
	if errBanner != nil {
		c.JSON(400, map[string]interface{}{
			"status": false,
			"errors": errBanner.Error(),
		})
		return
	}

	c.JSON(200, map[string]interface{}{
		"status": true,
		"data":   banners,
	})

	return

}
