package utils

import (
	"testing"
)

func TestUtilsRandomString(t *testing.T) {
	str := RandomStr(10)
	if len(str) == 0 || len(str) != 10 {
		t.Error("random string error")
	}
	t.Log(str)
}
