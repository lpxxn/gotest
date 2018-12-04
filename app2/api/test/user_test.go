package test

import (
	"testing"
)

func TestNewUserInfo(t *testing.T) {
	u := NewUserInfo()
	if len(u.Name) == 0 {
		t.Error("name is empty")
	}
}
