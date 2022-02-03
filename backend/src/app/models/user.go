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

func FindUser(id int) (*entity.User, error) {
	db := db.GetDB()
	var user entity.User
	println(id)
	if err := db.First(&user, id).Error; err != nil {
		return &user, err
	}
	return &user, nil
}

func FindAllUser() (*[]entity.User, error) {
	db := db.GetDB()
	var users []entity.User
	if err := db.Limit(3).Find(&users).Error; err != nil {
		return &users, err
	}
	return &users, nil
}