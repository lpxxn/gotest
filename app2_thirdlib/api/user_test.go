package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jarcoal/httpmock"
	"github.com/lpxxn/gotest/app2_thirdlib/model"
	"github.com/lpxxn/gotest/app2_thirdlib/testutils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestNewUserInfo(t *testing.T) {
	a := assert.New(t)

	router := gin.New()
	const path = "/newUserInfo"
	router.POST(path, NewUserInfo)

	body := url.Values{}
	body.Set("say", "hello world")
	rr, err := testutils.PostFormRequst(path + "?name=lp", router, body)
	a.Nil(err)

	user := &model.UserInfo{}
	err = json.Unmarshal(rr.Body.Bytes(), user)
	a.Nil(err)
	a.NotEqual(user.Name, "")
	a.NotEqual(user.Age, 0)
	t.Logf("%#v\n", user)
}

func TestUserInfoList(t *testing.T) {
	a := assert.New(t)

	router := gin.New()
	const path = "/userInfoList"
	router.GET(path, UserInfoList)

	rr, err := testutils.GetRequst(path, router)
	a.Nil(err)
	users := make([]model.UserInfo, 0)

	err = json.Unmarshal(rr.Body.Bytes(), &users)
	a.Nil(err)
	a.Len(users, 10)
}

func TestUserRoleList(t *testing.T) {
	a := assert.New(t)
	// mock http
	testutils.HttpMockActivateNonDefault()
	httpmock.RegisterNoResponder(
		httpmock.NewStringResponder(http.StatusOK, fmt.Sprintf(`
		[
		  {
			"id": 1,
			"name": "a"
		  },
		  {
			"id": 2,
			"name": "b"
		  },
		  {
			"id": 3,
			"name": "c"
		  }
		]
	`)))
	defer httpmock.DeactivateAndReset()

	router := gin.New()
	const path = "/userRoleList"
	router.GET(path, UserRoleList)
	rr, err := testutils.GetRequst(path, router)
	a.Nil(err)
	a.Equal(rr.Result().StatusCode, http.StatusOK)

	roles := make([]model.UserRole, 0)
	err = json.Unmarshal(rr.Body.Bytes(), &roles)
	a.Nil(err)
	a.NotEqual(len(roles), 0)
	t.Logf("len of roles: %d\n", len(roles))
}