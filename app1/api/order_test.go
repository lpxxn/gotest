package api

import "testing"

func TestGetOrderList(t *testing.T) {
	orderList := GetOrderList()
	if len(orderList) == 0 {
		t.Error("orderList is empty")
	}
	t.Logf("orderList len is %d\n", len(orderList))
}
