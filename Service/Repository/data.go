package repository

import "service/model"

type GetData interface {
	GetUserById(id int)(*model.User,error)
}

func NewGetData() GetData {
	return &getData{}
}

type getData struct {
}

func (g *getData) GetUserById(id int) (*model.User, error) {
	return &model.User{Name: "zhangsan", Age: 20}, nil
}
