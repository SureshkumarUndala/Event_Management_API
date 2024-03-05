package models

import (
	"errors"
	"time"

	"github.com/SureshkumarUndala/Event_Management_API/db"
	"github.com/SureshkumarUndala/Event_Management_API/middlewares"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user *User) Create() (int64, error) {

	row, err := db.DB.Exec("INSERT INTO users(Name, Email,Password) values(?,?,?)", user.Name, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	id, _ := row.LastInsertId()

	return id, nil

}

func (user *User) IsValidcredentials() error {
	var retrivedPassword string
	err := db.DB.QueryRow("SELECT id, password from users where email = ?", user.Email).Scan(&user.Id, &retrivedPassword)

	if err != nil {
		middlewares.Logger.Error().Printf("Error occured while authenticating the user %d", time.Now().UnixNano())
		return errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(retrivedPassword), []byte(user.Password))
	if err != nil {
		middlewares.Logger.Error().Printf("Error occured while validating user password %d", time.Now().UnixNano())
		return errors.New("invalid credentials")

	}

	return nil
}
