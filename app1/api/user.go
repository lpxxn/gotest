package api

import (
	"encoding/json"
	"fmt"
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

func HandleNewUser(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Printf("url parameter user name is %s\n", name)

	say := r.FormValue("say")
	fmt.Printf("req say:' %s '\n", say)
	newUser := NewUserInfo()
	jData, _ := json.Marshal(newUser)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}
