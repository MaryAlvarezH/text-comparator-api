package entity

import "time"

// User represents an application user.
type User struct {
	ID        int        `json:"id"`
	Name      string     `json:"name" binding:"required"`
	Lastname  string     `json:"lastname" binding:"required"`
	Email     string     `json:"email" gorm:"unique" binding:"required"`
	Password  string     `json:"password" binding:"required"`
	CreatedAt *time.Time `json:"createdAt"`
	DeleteAt  *time.Time `json:"deleteAt"`
}

func (User) TableName() string {
	return "app.users"
}
