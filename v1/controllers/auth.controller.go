package controllers

import (
	"api-goklaster/util"
	"api-goklaster/v1/models"
	"api-goklaster/v1/repositories"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req models.AuthLogin
	c.Bind(&req)
	errVal := req.Validate()
	if errVal != nil {
		c.JSON(400, map[string]interface{}{
			"status": false,
			"errors": errVal,
		})
		return
	}

	var result, errCheck = repositories.LoginCheck(req)
	if errCheck != nil {
		c.JSON(500, map[string]interface{}{
			"status": false,
			"errors": errCheck.Error(),
		})
		return
	}

	var token, errToken = util.GetToken(result)
	if errToken != nil {
		c.JSON(500, map[string]interface{}{
			"status": false,
			"errors": errToken,
		})
		return
	}

	c.JSON(200, map[string]interface{}{
		"status":  true,
		"message": "Token created",
		"token":   token,
	})

}
