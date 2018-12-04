package api

import "github.com/lpxxn/gotest/app1/utils"

type UserInfo struct {
	Name string
	Age  int
}

func (u *UserInfo) GetName() string {
	return u.Name
}

func (u *UserInfo) SetName(name string) {
	u.Name = name
}

func NewUserInfo() *UserInfo {
	return &UserInfo{Name: utils.RandomStr(3), Age: utils.RandomInt(5, 30)}
}
