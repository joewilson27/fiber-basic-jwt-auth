package repository

import (
	"errors"

	"github.com/joewilson27/fiber-basic-jwt-auth/models"
)

// Simulate a database call
func FindByCredentials(email, password string) (*models.User, error) {
	// Here you would query your database for the user with the given email
	if email == "test@mail.com" && password == "test1234567" {
		return &models.User{
			ID:             1,
			Email:          email,
			Password:       password,
			FavoritePhrase: "Hello, Mordor!",
		}, nil
	}
	return nil, errors.New("user not found")
}
