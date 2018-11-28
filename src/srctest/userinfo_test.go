package srctest

import (
	"github.com/lpxxn/gotest/src"
	"testing"
)

func TestNewUserInfo(t *testing.T) {
	u := src.NewUserInfo()
	if len(u.Name) == 0 {
		t.Error("name is empty")
	}
}