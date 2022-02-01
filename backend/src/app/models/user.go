package model

import (
		"app/models/entity"
		"app/db"
		)

func CreateUser(user *entity.User) (*entity.User, error) {
	db := db.GetDB()
	if err := db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}