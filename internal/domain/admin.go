package domain

type Admin struct {
	ID     uint `json:"id" gorm:"primaryKey;not null"`
	UserID uint `json:"user_id" gorm:"uniqueIndex;not null"`
	User   User

	Name    string `json:"name" gorm:"not null"`
	Phone   string `json:"phone" gorm:"not null"`
	Address string `json:"address" gorm:"not null"`
}
