package dto

type RegisterAdminProfileRequest struct {
	Name    string `json:"name" bindning:"required"`
	Phone   string `json:"phone" binding:"required"`
	Address string `json:"address"`
	Photo   string `json:"photo"`
}
