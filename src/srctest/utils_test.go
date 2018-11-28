package srctest

import (
	"github.com/lpxxn/gotest/src"
	"testing"
)


func TestUtilsRandomString(t *testing.T) {
	str := src.RandomStr(10)
	if len(str) == 0 || len(str) != 10 {
		t.Error("random string error")
	}
	t.Log(str)
}
