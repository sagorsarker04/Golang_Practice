package user_service

import (
	repository "service/Repository"
	"service/model"
)

type UserService interface {
	GetUser(id int) (*model.User, error)
	//GetUser2(id int)(*model.User,error)
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
