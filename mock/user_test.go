package d1

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_data "github.com/lpxxn/gotest/mock/dao"
)

func TestUser1(t *testing.T) {
	mockUser := mock_data.NewMockUser(gomock.NewController(t))
	mockUser.EXPECT().V(1, "2").Do(func(id int, name string) {
		t.Log(id, " ", name)
	}).Do(func(id int, name string) {
		t.Log("do2 ", id)
	}).Return("a", nil)
	// error 要把所有的参数都加上才可以，也不能多值
	//Do(func(name string) {
	//	t.Log("do2 ", name)
	//})
	//mockUser.EXPECT().V(1, "2").Return("a", nil)
	var u User = mockUser
	a, err := u.V(1, "2")
	t.Log(a, err)
}

func TestUser2(t *testing.T) {
	mockUser := mock_data.NewMockUser(gomock.NewController(t))
	mockUser.EXPECT().V(1, "2").AnyTimes()
	// 没有 return 会返回空
	//mockUser.EXPECT().V(1, "2").Return("a", nil)
	var u User = mockUser
	a, err := u.V(1, "2")
	t.Log(a, err)
}

func TestUser3(t *testing.T) {
	mockUser := mock_data.NewMockUser(gomock.NewController(t))
	mockUser.EXPECT().V(1, "2").DoAndReturn(func(idx int, n string) (string, error) {
		t.Log(idx, "  ", n)
		return "1", nil
	}).DoAndReturn(func(idx int, n string) (string, error) {
		t.Log(idx, "  ", n)
		return "2", nil
	})
	// 没有有多个 return 会返回最后一个
	var u User = mockUser
	a, err := u.V(1, "2")
	t.Log(a, err)
	mockUser.EXPECT().V(1, "2").Return("c", nil).Return("d", nil)
	a, err = u.V(1, "2")
	t.Log(a, err)
}
