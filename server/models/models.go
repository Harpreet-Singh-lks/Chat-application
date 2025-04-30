package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey;AutoIncrement"`
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password []byte `json:"-"`
}
