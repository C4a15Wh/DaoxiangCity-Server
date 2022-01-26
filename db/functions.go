package db

import "dxcserver/model"

func SelectUserById(id int) (model.User, error) {
	var UserInfo model.User
	err := exeDB.Where(`uid = ?`, id).Find(&UserInfo).Error
	return UserInfo, err
}

func SelectUserByName(nickname string) (model.User, error) {
	var UserInfo model.User
	err := exeDB.Where(`nickname = ?`, nickname).First(&UserInfo).Error
	return UserInfo, err
}
