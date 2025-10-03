package models

type User struct {
	ID          int64  `json:"id" gorm:"primaryKey,autoIncrement"`
	Email       string `json:"email" gorm:"unique"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
