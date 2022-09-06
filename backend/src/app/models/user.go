package model

import (
	"app/db"
	"app/models/entity"
	"app/parameters"
)

func CreateUser(req parameters.PostUserParams) (*entity.User, error) {
	db := db.GetDB()
	user := &entity.User{
		DisplayName: req.DisplayName,
	}
	if err := db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func FindUser(id uint64) (*entity.User, error) {
	db := db.GetDB()
	var user entity.User
	if err := db.First(&user, id).Error; err != nil {
		return &user, err
	}
	return &user, nil
}

func FindAllUser() ([]*entity.User, error) {
	db := db.GetDB()
	var users []*entity.User
	if err := db.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}
