package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id uint32 `json:"id"`
	// Authentication Data
	Username string `gorm:"primary_key;auto_increment" json:"username"`
	Email    string `gorm:"size:255;not null;unique" json:"author"`
	Password string `gorm:"size:100;not null;" json:"password"`

	// Backend Data
	// -- Account Details
	CreatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdat"`
	ConfirmedAccount bool      `gorm:"default:false" json:"confirmedaccount"`
	BlockedAccount   bool      `gorm:"default:false" json:"blockedaccount"`
	UpdatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedat"`
	// -- Providers
	ProviderFacebookId string `gorm:"size:255;unique" json:"providerfacebookid"`
	ProviderGoogleId   string `gorm:"size:255;unique" json:"providergoogleid"`
	ProviderTwitterId  string `gorm:"size:255;unique" json:"providertwitterid"`
	// -- Tokens
	AccountConfirmationToken string `gorm:"size:255;unique" json:"accountconfirmationtoken"`
	ResetPasswordToken       string `gorm:"size:255;unique" json:"resetpasswordtoken"`

	// Personal Data
	FirstName      string `gorm:"size:100" json:"firstname"`
	LastName       string `gorm:"size:100" json:"lastname"`
	ProfilePicture string `gorm:"size:255" json:"profilepicture"`
}

// Helpers

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.Id = 0
	u.FirstName = html.EscapeString(strings.TrimSpace(u.FirstName))
	u.LastName = html.EscapeString(strings.TrimSpace(u.LastName))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Password == "" {
			return errors.New("palavra-passe é obrigatória")
		}
		if u.Email == "" {
			return errors.New("email é obrigatório")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("email inválido")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("palavra-passe é obrigatória")
		}
		if u.Email == "" {
			return errors.New("email é obrigatório")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("email inválido")
		}
		return nil

	default:
		if u.Username == "" {
			return errors.New("nome de utilizador é obrigatório")
		}
		if u.Password == "" {
			return errors.New("palavra-passe é obrigatória")
		}
		if u.Email == "" {
			return errors.New("email é obrigatório")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("email inválido")
		}
		return nil
	}
}
