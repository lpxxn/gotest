package api

import (
	"github.com/lpxxn/gotest/app1/model"
	"github.com/lpxxn/gotest/app1/utils"
)

func NewUserInfo() *model.UserInfo {
	return &model.UserInfo{Name: utils.RandomStr(3), Age: utils.RandomInt(5, 30)}
}

func UserInfoList() []*model.UserInfo {
	userInfoList := make([]*model.UserInfo, 0)
	for i := 0; i < 10; i++ {
		userInfoList = append(userInfoList, NewUserInfo())
	}
	return userInfoList
}
