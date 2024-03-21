package models

type User struct {
	ID       int    `json:"id,omitempty" gorm:"primaryKey"`
	Name     string `json:"name,omitempty"`
	Age      int    `json:"age,omitempty"`
	Address  string `json:"address,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}