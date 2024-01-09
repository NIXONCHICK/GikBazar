package model

type User struct {
	ID uint `json:"User_id" gorm:"primaryKey"`

	Login    string `json:"User_login" gorm:"not null;column:login;size:255"`
	Password string `json:"Password" gorm:"not null;column:password;size:255"`
	Role     string `json:"Role" gorm:"not null;column:role;size:255"`
}
