package models

type UserModel struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func (UserModel) TableName() string {
	return "user"
}
