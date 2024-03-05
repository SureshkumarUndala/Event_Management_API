package controllers

import (
	"net/http"
	"time"

	"github.com/SureshkumarUndala/Event_Management_API/api_test/utils"
	"github.com/SureshkumarUndala/Event_Management_API/db"
	"github.com/SureshkumarUndala/Event_Management_API/middlewares"
	"github.com/SureshkumarUndala/Event_Management_API/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func encryptPasssword(textpassword string) (string, error) {
	hashedpass, err := bcrypt.GenerateFromPassword([]byte(textpassword), 10)
	if err != nil {
		return "", err
	}
	return string(hashedpass), nil

}

func Login(c *gin.Context) {

	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		middlewares.Logger.Error().Printf("Error occured while parsing the user data from the login  request %d", time.Now().UnixNano())
		c.JSON(400, gin.H{"status": "400", "message": "All fields are required"})
		return
	}
	err = user.IsValidcredentials()

	if err != nil {
		middlewares.Logger.Error().Printf("Error occured while validating user credentials for login request %d", time.Now().UnixNano())
		c.JSON(400, gin.H{"status": "Invalid credentials"})
		return

	}
	token, err := utils.CreateJWT(user.Email)
	if err != nil {
		middlewares.Logger.Error().Printf("Error occured while generating JWT token %d", time.Now().UnixNano())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "401", "message": err.Error()})
		return

	}
	middlewares.Logger.Info().Printf("user login request successfull %d", time.Now().UnixNano())
	c.JSON(200, gin.H{"status": "200", "message": "User login successful", "token": token})

}

func RegisterUser(c *gin.Context) {
	middlewares.Logger.Info().Printf("RegisterUser controller was called %d", time.Now().UnixNano())
	var Newuser models.User
	err := c.ShouldBindJSON(&Newuser)

	if err != nil {
		middlewares.Logger.Error().Printf("Error occured while parsing the user data from the signup request %d", time.Now().UnixNano())

		c.JSON(400, gin.H{"status": "400", "message": "All fields are required"})
		return
	}
	var count int

	err = db.DB.QueryRow("select COUNT(id) from users where email = ?", Newuser.Email).Scan(&count)

	if err != nil {
		middlewares.Logger.Error().Printf("Error occured while query the user from database %d", time.Now().UnixNano())
		c.JSON(400, gin.H{"status": "400", "message": err.Error()})
	}

	if count > 0 {
		middlewares.Logger.Error().Printf("Error occured while registered user try to create another account %d", time.Now().UnixNano())
		c.JSON(400, gin.H{"status": "400", "message": "user already had an account"})
		return
	}

	hashedPassword, err := encryptPasssword(Newuser.Password)
	if err != nil {
		middlewares.Logger.Error().Printf("Error occured while generating hashed password %d", time.Now().UnixNano())

	}
	Newuser.Password = hashedPassword
	_, err = Newuser.Create()
	if err != nil {
		middlewares.Logger.Error().Printf("Error occured while saving user details in database %d", time.Now().UnixNano())
		c.JSON(400, gin.H{"status": "400", "message": err})
		return

	}
	middlewares.Logger.Info().Printf("user Registration successful %d", time.Now().UnixNano())

	c.JSON(200, gin.H{"status": "200", "message": "user registered successfully"})
}

func Forgotpassword(c *gin.Context) {
	c.JSON(200, gin.H{"status": "200", "message": "forgotpassword link sent successfully"})

}

func Resetpassword(c *gin.Context) {
	c.JSON(200, gin.H{"status": "200", "message": "Password Reset successfully link sent successfully"})

}
