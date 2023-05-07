package lib

import (
	"mini_project/database"
	"mini_project/middleware"
	"mini_project/model"
	"mini_project/util"
)

func LoginUsers(user *model.User) (interface{}, error) {
	var err error
	var userDB model.User
	if err = database.DB.First(&userDB, "email = ?", user.Email).Error; err != nil {
		return nil, err
	}

	if err = util.ComparePassword(userDB.Password, user.Password); err != nil {
		return nil, err
	}

	Token, err := middleware.CreateToken(userDB.ID)
	if err != nil {
		return nil, err
	}

	return Token, nil
}
