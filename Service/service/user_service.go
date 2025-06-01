package user_service

import (
	repository "service/Repository"
	"service/model"
)

type UserService interface {
	GetUser(id int) (*model.User, error)
	//GetUser2(id int)(*model.User,error)
	CreateUser(user model.User) (model.User, error)
	CreateTable() error
}

type userService struct {
	//Getuser()
	data repository.GetData
}

func NewUserService(d repository.GetData) UserService {
	return &userService{data: d}
}

// type xyz (int){
// 	fmt.Println("string")
// }

func (u *userService) GetUser(id int) (*model.User, error) {
	user, err := u.data.GetUserById(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) CreateUser(user model.User) (model.User, error) {
	err := u.data.CreateUser(user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u *userService) CreateTable() error {
	err := u.data.CreateTable()
	if err != nil {
		return err
	}
	return nil
}
