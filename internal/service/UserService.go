package service

import (
	"TestDemo/internal/model"
	"errors"
	"fmt"
)

var UserMap map[string]model.Users = make(map[string]model.Users)

func InsertUsers(users model.Users) (map[string]model.Users, error) {
	if ids := users.GetId(); ids == "" {
		return UserMap, errors.New("没有主键未插入进去")
	} else {
		UserMap[users.GetId()] = users
		return UserMap, nil
	}

}

func DeleteUsers(id string) (map[string]model.Users, error) {

	var usr model.Users
	if u := UserMap[id]; u == usr {
		fmt.Println("不存在该用户，请删除存在的用户")
		return UserMap, errors.New("不存在该用户，请删除存在的用户,报错")
	} else {
		delete(UserMap, id)
		return UserMap, nil
	}

}
