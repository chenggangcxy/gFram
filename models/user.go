package models

//import "package"

type User struct {
	Id       int
	Username string
	Age      int
	Email    string
	AddTime  int
	Iconurl  string
}

func (User) TableName() string {
	return "user"
}
