package dto

import "time"

// STEP1
type RegisterRequest struct {
	Username string `json:"username" binding:"required, min=2, max=100"`
	Password string `json:"password" binding:"required, min=8"`
}

// STEP2
type RegisterAdminProfileRequest struct {
	Name    string `json:"name" bindning:"required"`
	Phone   string `json:"phone" binding:"required"`
	Address string `json:"address"`
	Photo   string `json:"photo"`
}

type UserResponse struct {
	ID        uint64    `json:"id"`
	Username  string    `json:"username"`
	Roles     []string  `json:"roles"`
	CreatedAt time.Time `json:"created_at"`
}
