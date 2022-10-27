package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	// Authentication Data
	Username string `gorm:"size:255;unique" json:"username"`
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:100;not null;" json:"password"`

	// Backend Data
	// -- Account Details
	CreatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdat"`
	ConfirmedAccount bool      `gorm:"default:false" json:"confirmedaccount"`
	BlockedAccount   bool      `gorm:"default:false" json:"blockedaccount"`
	Role             string    `gorm:"size:255;default:'client'" json:"role"`
	UpdatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedat"`
	// -- Providers
	ProviderFacebookId string `gorm:"size:255" json:"providerfacebookid"`
	ProviderGoogleId   string `gorm:"size:255" json:"providergoogleid"`
	ProviderTwitterId  string `gorm:"size:255" json:"providertwitterid"`
	// -- Tokens
	AccountConfirmationToken string `gorm:"size:255" json:"accountconfirmationtoken"`
	ResetPasswordToken       string `gorm:"size:255" json:"resetpasswordtoken"`

	// Personal Data
	FirstName      string `gorm:"size:100" json:"firstname"`
	LastName       string `gorm:"size:100" json:"lastname"`
	ProfilePicture string `gorm:"size:255" json:"profilepicture"`
}

type LoginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type userResponse struct {
	Username  string    `json:"username"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

type UserResponse struct {
	Id             uuid.UUID
	Username       string
	Email          string
	FirstName      string
	LastName       string
	ProfilePicture string
	CreatedAt      time.Time
	Role           string
}

func NewUserResponse(user User) userResponse {
	return userResponse{
		Username:  user.Username,
		FullName:  user.FirstName + " " + user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

// HashPassword returns the bcrypt hash of the password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

// CheckPassword checks if the provided password is correct or not
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
