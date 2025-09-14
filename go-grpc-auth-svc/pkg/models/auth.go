package models

type User struct {
	Id       int64  `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `gorm:"default:user" json:"role"` // "user" or "admin"
}
