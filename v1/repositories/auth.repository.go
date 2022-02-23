package repositories

import (
	"api-goklaster/config"
	"api-goklaster/entities"
	"api-goklaster/util"
	"api-goklaster/v1/models"
	"fmt"
)

var db = config.InitDB().Debug()

func LoginCheck(post models.AuthLogin) (util.DataJwt, error) {
	var user entities.User
	var dataJwt = util.DataJwt{}
	query := db.
		Select("username, permission").
		Where("username = ? AND password = md5(?)",
			post.Username,
			post.Password).
		Find(&user)

	if query.Error != nil {
		return dataJwt, query.Error
	}

	if user.Username == "" {
		return dataJwt, fmt.Errorf("username or password is wrong")
	}

	dataJwt.Username = user.Username
	dataJwt.Permission = user.Permission

	return dataJwt, nil
}
