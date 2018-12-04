package api

import (
	"testing"
)

func TestNewUserInfo(t *testing.T) {
	u := NewUserInfo()
	if len(u.Name) == 0 {
		t.Error("name is empty")
	}
}

func TestUserInfoList(t *testing.T) {
	userInfoList := UserInfoList()
	if len(userInfoList) != 10 {
		t.Error("userInfoList is empty")
	}
	t.Logf("userInfoList len := %d\n", len(userInfoList))
}