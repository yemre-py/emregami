package domain

// authoriation and authentication
type Auth struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"uniqueIndex" validate:"required, email"`
	Username string `json:"username" gorm:"uniqueIndex" validate:"required, min=3, max=32, username"`
	Password string `json:"password" gorm:"not null" validate:"required, password"`
}
