package domain 

type UserModel struct {
	Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
}


func (UserModel) TableName() string {
	return  "users"
}