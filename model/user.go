package model

type User struct {
	ID       uint   `gorm:"unique;primaryKey;autoIncrement" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
