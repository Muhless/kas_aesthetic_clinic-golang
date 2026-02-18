package dto

import "time"

// STEP1
type RegisterRequest struct {
	Username string `json:"username" binding:"required, min=2, max=100"`
	Password string `json:"password" binding:"required, min=8"`
	Role     string `json:"role" binding:"required, oneof=admin dokter perawat"`
}

type UserResponse struct {
	ID        uint64    `json:"id"`
	Username  string    `json:"username"`
	Roles     []string  `json:"roles"`
	CreatedAt time.Time `json:"created_at"`
}
