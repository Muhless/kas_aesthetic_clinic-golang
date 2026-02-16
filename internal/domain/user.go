package domain

import "time"

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`

	Roles []Role `json:"role" gorm:"many2many:user_roles"`

	CreatedAt time.Time `json:"created_at" gorm:""`
	UpdatedAt time.Time `json:"updated_at" gorm:""`
}
