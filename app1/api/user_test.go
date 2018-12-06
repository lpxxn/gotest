package api

import (
	"encoding/json"
	"github.com/lpxxn/gotest/app1/model"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
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

func BenchmarkUserInfoList(b *testing.B) {
	b.StopTimer()
	// do something
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		// pretend delay
		//time.Sleep(time.Millisecond * 500)
		userInfoList := UserInfoList()
		if len(userInfoList) != 10 {
			b.Error("userInfoList is empty")
		}
	}
}

func TestHandleNewUser(t *testing.T) {
	postBody := url.Values{}
	postBody.Add("say", "hello world")
	req := httptest.NewRequest(http.MethodPost, "http://localhost/createuser?name=linus", strings.NewReader(postBody.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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
