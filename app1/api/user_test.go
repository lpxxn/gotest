package api

import (
	"encoding/json"
	"github.com/lpxxn/gotest/app1/model"
	"net/http"
	"net/http/httptest"
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

func TestHandleNewUser(t *testing.T) {
	req, _ := http.NewRequest(http.MethodPost, "newUser", nil)
	w := httptest.NewRecorder()
	HandleNewUser(w, req)

	if w.Code != http.StatusOK {
		t.Error("new user api error")
	}
	if w.Body.Len() == 0 {
		t.Error(" response is empty")
	}
	user := &model.UserInfo{}
	err := json.Unmarshal(w.Body.Bytes(), user)
	if err != nil {
		t.Error("response data error")
	}
	t.Logf("create user api response : %#v", user)
}

