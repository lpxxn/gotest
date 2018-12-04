package api

import (
	"encoding/json"
	"github.com/lpxxn/gotest/app1/model"
	"github.com/lpxxn/gotest/app1/utils"
	"net/http"
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


func HandleNewUser(w http.ResponseWriter,r *http.Request) {
	newUser := NewUserInfo()
	jData, _ := json.Marshal(newUser)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}