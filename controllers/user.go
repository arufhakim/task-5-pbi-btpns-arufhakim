package controllers

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"task-5-pbi-btpns-arufhakim/app"
	"task-5-pbi-btpns-arufhakim/helpers"
	"task-5-pbi-btpns-arufhakim/middlewares"
	"task-5-pbi-btpns-arufhakim/models"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (uc *UserController) Update(c *gin.Context) {
	authHeader := c.Request.Header["Authorization"][0]
	jwtToken := strings.Split(authHeader, " ")[1]
	userIdFromJwtToken, err := middlewares.ExtractJWTToken(jwtToken)
	userId := int(userIdFromJwtToken.(float64))

	if err != nil {
		log.Println("JWT error: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	var userInput app.UserValidation

	if err := c.Bind(&userInput); err != nil {
		log.Println("Binding error: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	userIdFromParam, _ := strconv.Atoi(c.Param("userId"))

	if userId == userIdFromParam {
		emailRegistered, emailId, err := helpers.IsRegistered(userInput.Email)

		if err != nil {
			log.Println("Email error: ", err.Error())
			helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

			return
		}

		if emailRegistered && emailId != uint(userId) {
			helpers.SendResponse(c, http.StatusBadRequest, "User already exist", nil)

			return
		}

		userForUpdate := models.User{
			Username: userInput.Username,
			Email:    userInput.Email,
		}
		userForUpdate.ID = uint(userId)

		if err := user.Update(&userForUpdate); err != nil {
			log.Println("Update error: ", err)
			helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)
			return
		}

		helpers.SendResponse(c, http.StatusOK, "Successfully update user", nil)
		return
	}

	helpers.SendResponse(c, http.StatusUnauthorized, "Error not found", nil)
}

func (uc *UserController) Delete(c *gin.Context) {
	authHeader := c.Request.Header["Authorization"][0]
	jwtToken := strings.Split(authHeader, " ")[1]
	userIdFromJwtToken, err := middlewares.ExtractJWTToken(jwtToken)
	userId := int(userIdFromJwtToken.(float64))

	if err != nil {
		log.Println("JWT error: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	userIdFromParam, _ := strconv.Atoi(c.Param("userId"))

	if userId == userIdFromParam {
		if err := user.Delete(uint(userId)); err != nil {
			helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

			return
		}

		helpers.SendResponse(c, http.StatusOK, "Successfully delete user", nil)

		return
	}

	helpers.SendResponse(c, http.StatusUnauthorized, "Error not Found", nil)
}
