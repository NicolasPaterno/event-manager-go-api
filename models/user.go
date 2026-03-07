package models

import (
	"errors"

	"github.com/nicao/minimal-goapi/db"
	"github.com/nicao/minimal-goapi/utils"
)

type User struct {
	ID       int64  `json:"id,omitempty" example:"1"`
	Email    string `json:"email" binding:"required" example:"user@example.com"`
	Password string `json:"password" binding:"required" example:"senha123"`
}

// UserCredentials representa o body de login e signup (apenas email e password)
type UserCredentials struct {
	Email    string `json:"email" binding:"required" example:"user@example.com"`
	Password string `json:"password" binding:"required" example:"senha123"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email,password) VALUES (?,?)"
	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	HashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := statement.Exec(u.Email, HashPassword)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()

	u.ID = userId
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return errors.New("credentials invalid")
	}

	passwordIsValid := utils.CheckPassword(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}
