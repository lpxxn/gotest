package api

import (
	"encoding/json"
	"fmt"
	"github.com/lpxxn/gotest/app1/model"
	"net/http"
	"net/http/httptest"
	"runtime"
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
	n := runtime.GOMAXPROCS(-1)
	b.Log("cpu number : ", n)
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
	fmt.Println("N: ", b.N)
}

func TestHandleNewUser(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "http://localhost/createuser", nil)
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
