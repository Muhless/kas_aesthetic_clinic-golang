package domain

type Role struct {
	ID   uint64   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"uniqueIndex;not null"`

	Permissions []Permission `json:"permissions" gorm:"many2many:role_permissions;"`
}
