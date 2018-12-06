package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lpxxn/gotest/app2_thirdlib/model"
	"github.com/lpxxn/gotest/app2_thirdlib/utils"
	"io/ioutil"
	"net/http"
)

func NewUserInfo(c *gin.Context) {
	name := c.Query("name")
	fmt.Printf("name is ' %s '\n", name)
	say, _ := c.GetPostForm("say")
	fmt.Printf("say ' %s '\n", say)
	user := newUserInfo()
	c.JSON(http.StatusOK, user)
}

func UserInfoList(c *gin.Context) {
	userInfoList := make([]*model.UserInfo, 0)
	for i := 0; i < 10; i++ {
		userInfoList = append(userInfoList, newUserInfo())
	}
	c.JSON(http.StatusOK, userInfoList)
}

func UserRoleList(c *gin.Context) {
	roles, err := allUserRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusOK, roles)
}

func allUserRoles() ([]model.UserRole, error) {
	// 假如需要用到其他服务的接口
	res, err := utils.HttpClient.Get("http://localhost/alluserroles")
	if err != nil {
		return nil, err
	}
	roles := make([]model.UserRole, 0)
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &roles)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func newUserInfo() *model.UserInfo {
	return &model.UserInfo{Name: utils.RandomStr(3), Age: utils.RandomInt(5, 30)}
}
