package models

import "time"

type User struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreateAt  time.Time
	FirstName string `json:"first_name"`
	LastName  string `json:"Last_name"`
}
