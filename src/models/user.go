package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type AuthResponse struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}

// Fetches a user from the database
func FetchUser(id uint) (*User, error) {
	var user User
	err := Database.Where("id = ?", id).First(&user).Error
	if err != nil {
		return &User{}, err
	}
	return &user, nil
}

func FetchUserByEmail(email string) User {
	var userFromDb User
	Database.Where("email = ?", email).First(&userFromDb)

	return userFromDb
}

func (user *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashedPassword)
	return err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (user *User) Register() (*AuthResponse, error) {
	var err error
	userFromDb := FetchUserByEmail(user.Email)

	if userFromDb.Email != "" {
		err = errors.New("email already taken")
		return &AuthResponse{}, err
	}

	err = user.HashPassword()
	if err != nil {
		return &AuthResponse{}, err
	}

	err = Database.Model(&user).Create(user).Error
	if err != nil {
		return &AuthResponse{}, err
	}

	token, err := GenerateJWT(user.ID)
	if err != nil {
		return &AuthResponse{}, err
	}

	response := AuthResponse{
		User:  user,
		Token: token,
	}

	return &response, nil
}

func (user *User) Login() (*AuthResponse, error) {
	var err error
	userFromDb := FetchUserByEmail(user.Email)

	if userFromDb.Email == "" {
		err = errors.New("User or password incorrect")
		return &AuthResponse{}, err
	}

	var isCheckedPassword = CheckPasswordHash(user.Password, userFromDb.Password)
	if !isCheckedPassword {
		err = errors.New("User or password incorrect")
		return &AuthResponse{}, err
	}

	token, err := GenerateJWT(user.ID)
	if err != nil {
		return &AuthResponse{}, err
	}

	response := AuthResponse{
		User:  &userFromDb,
		Token: token,
	}

	return &response, nil
}

func (user *User) UpdateUser(id string) (*User, error) {
	if user.Password != "" {
		err := user.HashPassword()
		if err != nil {
			return &User{}, err
		}
	}

	err := Database.Model(&User{}).Where("id = ?", id).Updates(user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}
