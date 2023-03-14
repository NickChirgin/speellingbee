package types

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
	HashedPassword string
}

func NewUser(username, password string) (*User, error) {
	hashedPW, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Cannot hash password: %w", err)
	}

	user := &User{
		Username: username,
		HashedPassword: string(hashedPW),
	}
	return user, nil
}

func (u *User) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password))
	return err == nil
}