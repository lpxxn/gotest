package api

import "github.com/lpxxn/gotest/app1/model"

func GetOrderList() model.OrderInfoList {
	oil := model.OrderInfoList{}
	for i := 0; i < 10; i++ {
		oil = append(oil, model.OrderById(i))
	}
	return oil
}
