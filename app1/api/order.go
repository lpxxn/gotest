package api

import "github.com/lpxxn/gotest/app1/utils"

type OrderInfo struct {
	Id         int
	TotalMoney int
}

func OrderById(id int) OrderInfo {
	return OrderInfo{Id: id, TotalMoney: utils.RandomInt(10, 10000)}
}

type OrderInfoList []OrderInfo

func GetOrderList() OrderInfoList {
	oil := OrderInfoList{}
	for i := 0; i < 10; i++ {
		oil = append(oil, OrderById(i))
	}
	return oil
}
