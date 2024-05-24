package model

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var SigningKey = []byte("secret_key")

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"` // gorm:autoCreateTime
	UpdatedAt time.Time `json:"updated_at"` // gorm:autoUpdateTime
	Baskets   []Basket  `json:"baskets,omitempty" gorm:"foreignKey:UserID;references:ID"`
}

func CreateUser(db *gorm.DB, username, password string) (*User, error) {
	var hashedPassword, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	var user = &User{
		Username: username,
		Password: string(hashedPassword),
	}
	var result = db.Create(user)
	if result.Error != nil {
		return nil, UserCreateError{errors.New("failed to create user")}
	}
	return user, nil
}

func LoginUser(db *gorm.DB, username, password string) (string, error) {
	var user User
	var result = db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return "", UsernamePasswordMismatchError{errors.New("invalid username or password")}
		}
		return "", UserRetreiveError{errors.New("failed to get user")}
	}
	var err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", UsernamePasswordMismatchError{errors.New("invalid username or password")}
	}
	var claims = jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	}
	var token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(SigningKey)
	if err != nil {
		return "", TokenGenerationError{errors.New("failed to generate token")}
	}
	return signedToken, nil
}

func DeleteUser(db *gorm.DB, userID int64) error {
	var user User
	var result = db.First(&user, userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return UserNotFoundError{errors.New("user not found")}
		}
		return UserRetreiveError{errors.New("failed to get user")}
	}
	result = db.Delete(&user)
	if result.Error != nil {
		return UserDeleteError{errors.New("failed to delete user")}
	}
	return nil
}
