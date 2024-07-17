package models

import (
	"html"
	"strings"
	"time"

	"backend-tugas-reactjs/utils/token"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User struct with book relationship
type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Username  string    `gorm:"not null;unique" json:"username"`
	Email     string    `gorm:"not null;unique" json:"email"`
	Password  string    `gorm:"not null;" json:"password"`
	Books     []Book    `json:"-" gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// VerifyPassword compares the hashed password with the provided password
func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// LoginCheck verifies the username and password and generates a token
func LoginCheck(username string, password string, db *gorm.DB) (string, error) {
	var err error
	u := User{}

	err = db.Model(User{}).Where("username = ?", username).Take(&u).Error
	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

// SaveUser hashes the password and saves the user in the database
func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	// Turn password into hash
	hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if errPassword != nil {
		return &User{}, errPassword
	}
	u.Password = string(hashedPassword)
	// Remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	var err error = db.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}
