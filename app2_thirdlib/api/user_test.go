package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lpxxn/gotest/app2_thirdlib/model"
	"github.com/lpxxn/gotest/app2_thirdlib/testutils"
	"github.com/stretchr/testify/assert"
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