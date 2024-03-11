package models

import "time"

type User struct {
	Id       int       `json:"id" gorm:"primaryKey, autoIncrement"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Date     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"date"`
}
