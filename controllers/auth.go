package controllers

import (
	"log"
	"net/http"
	"strconv"
	"task-5-pbi-btpns-arufhakim/app"
	"task-5-pbi-btpns-arufhakim/controllers/queries"
	"task-5-pbi-btpns-arufhakim/helpers"
	"task-5-pbi-btpns-arufhakim/middlewares"
	"task-5-pbi-btpns-arufhakim/models"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

var user queries.UserQuery

func (ua *AuthController) Register(c *gin.Context) {
	var userInput app.UserValidation

	if err := c.Bind(&userInput); err != nil {
		log.Println("Binding error: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	isValidated, err := helpers.ValidateUserInputForAuthentication(userInput)

	if err != nil || !isValidated {
		log.Println("Validation error: " + err.Error())
		helpers.SendResponse(c, http.StatusBadRequest, err.Error(), nil)

		return
	}

	emailRegistered, _, err := helpers.IsRegistered(userInput.Email)

	if err != nil {
		log.Println("Email error: ", err.Error())
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	if emailRegistered {
		helpers.SendResponse(c, http.StatusBadRequest, "User already exist", nil)

		return
	}

	hashedPassword, err := helpers.HashPassword(userInput.Password)

	if err != nil {
		log.Println("Hashing error: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, "Error while registering account", nil)

		return
	}

	userForDatabase := models.User{
		Username: userInput.Username,
		Email:    userInput.Email,
		Password: hashedPassword,
	}

	if err = user.Save(&userForDatabase); err != nil {
		log.Println("Save error: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, "Error while registering account", nil)

		return
	}

	helpers.SendResponse(c, http.StatusCreated, "Successfully registered", nil)
}

func (ua *AuthController) Login(c *gin.Context) {
	var userInput app.LoginValidation

	if err := c.Bind(&userInput); err != nil {
		log.Println("Binding error: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	isValidated, err := helpers.ValidateUserInputForAuthentication(userInput)

	if err != nil || !isValidated {
		log.Println("Validation error: " + err.Error())
		helpers.SendResponse(c, http.StatusBadRequest, err.Error(), nil)

		return
	}

	emailRegistered, _, err := helpers.IsRegistered(userInput.Email)

	if err != nil {
		log.Println("Email error: ", err.Error())
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	if !emailRegistered {
		helpers.SendResponse(c, http.StatusUnauthorized, "Wrong password or email", nil)

		return
	}

	userData, err := user.Get(userInput.Email)

	if err != nil {
		log.Println("Hashing error: ", err.Error())
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	if err := helpers.ComparePassword(userInput.Password, userData.Password); err != nil {
		helpers.SendResponse(c, http.StatusUnauthorized, "Wrong password or email", nil)

		return
	}

	jwtToken, err := middlewares.CreateJWTToken(userData.ID)

	if err != nil {
		log.Println("JWT error: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	data := map[string]string{
		"token":   jwtToken,
		"user_id": strconv.Itoa(int(userData.ID)),
	}

	helpers.SendResponse(c, http.StatusOK, "Successfully Login", data)
}
