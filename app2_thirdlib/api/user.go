package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lpxxn/gotest/app2_thirdlib/model"
	"github.com/lpxxn/gotest/app2_thirdlib/utils"
	"net/http"
)

func NewUserInfo(c *gin.Context) {
	user := newUserInfo()
	c.JSON(http.StatusOK, user)
}

func UserInfoList(c *gin.Context) {
	userInfoList := make([]*model.UserInfo, 0)
	for i := 0; i < 10; i++ {
		userInfoList = append(userInfoList, newUserInfo())
	}
}

func newUserInfo() *model.UserInfo {
	return &model.UserInfo{Name: utils.RandomStr(3), Age: utils.RandomInt(5, 30)}
}
