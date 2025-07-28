package models

type User struct {
	Id    int
	Name  string
	Email string
}

func NewUser(id int, name, email string) *User {
	return &User{
		Id:    id,
		Name:  name,
		Email: email,
	}
}
