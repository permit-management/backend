package domain

type UserModel struct {

	UserID      string `json:"user_id"`
	Username    string `json:"name" gorm:"column:name"`               
	NumberPhone string `json:"phone_number" gorm:"column:number_phone"`
	Email       string `json:"email"`
	Password    string `json:"-"`
}

func (UserModel) TableName() string {
	return "tbl_users"
}
