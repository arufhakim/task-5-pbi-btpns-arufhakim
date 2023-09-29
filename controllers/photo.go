package controllers

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"task-5-pbi-btpns-arufhakim/controllers/queries"
	"task-5-pbi-btpns-arufhakim/controllers/response"
	"task-5-pbi-btpns-arufhakim/helpers"
	"task-5-pbi-btpns-arufhakim/middlewares"
	"task-5-pbi-btpns-arufhakim/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PhotoController struct{}

var photoQuery queries.PhotoQuery

func (p *PhotoController) Create(c *gin.Context) {
	authHeader := c.Request.Header["Authorization"]

	if authHeader == nil {
		log.Println("Header required")
		helpers.SendResponse(c, http.StatusBadRequest, "Header required", nil)

		return
	}

	authorization := authHeader[0]

	if authorization == "" {
		log.Println("Token required")
		helpers.SendResponse(c, http.StatusBadRequest, "Token required", nil)

		return
	}

	jwtToken := strings.Split(authorization, " ")[1]
	userIdFromJwtToken, err := middlewares.ExtractJWTToken(jwtToken)
	userId := int(userIdFromJwtToken.(float64))

	if err != nil {
		log.Println("JWT error: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var photo models.Photo
	if err := c.Bind(&photo); err != nil {
		log.Println("Binding error: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	file, err := c.FormFile("file")

	if err != nil {
		log.Println("Binding error: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	filePath, err := helpers.SaveFileToDir(c, file)

	if err != nil {
		log.Println("Save error: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	photo.UserID = uint(userId)
	photo.PhotoURL = filePath

	if err := photoQuery.Save(&photo); err != nil {
		log.Println("Save error: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	data := response.Photo{
		ID:       photo.ID,
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoURL: photo.PhotoURL,
		UserID:   photo.UserID,
	}

	helpers.SendResponse(c, http.StatusCreated, "Successfully upload photo", data)
}

func (p *PhotoController) Get(c *gin.Context) {
	authHeader := c.Request.Header["Authorization"]

	if authHeader == nil {
		log.Println("Header required")
		helpers.SendResponse(c, http.StatusBadRequest, "Header required", nil)

		return
	}

	authorization := authHeader[0]

	if authorization == "" {
		log.Println("Token required")
		helpers.SendResponse(c, http.StatusBadRequest, "Token required", nil)

		return
	}

	jwtToken := strings.Split(authorization, " ")[1]
	userIdFromJwtToken, err := middlewares.ExtractJWTToken(jwtToken)
	userId := int(userIdFromJwtToken.(float64))

	if err != nil {
		log.Println("JWT error: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	photo, err := photoQuery.Get(uint(userId))

	if err == gorm.ErrRecordNotFound {
		log.Println("Error not found: ")
		helpers.SendResponse(c, http.StatusNotFound, "Error not found", nil)

		return
	} else if err != nil {
		log.Println("Error not found: ", err)
		helpers.SendResponse(c, http.StatusNotFound, err.Error(), nil)

		return
	}

	data := response.Photo{
		ID:       photo.ID,
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoURL: photo.PhotoURL,
		UserID:   photo.UserID,
	}

	helpers.SendResponse(c, http.StatusOK, "Successfully get photo", data)
}

func (p *PhotoController) Update(c *gin.Context) {
	authHeader := c.Request.Header["Authorization"]

	if authHeader == nil {
		log.Println("Header required")
		helpers.SendResponse(c, http.StatusBadRequest, "Header required", nil)

		return
	}

	authorization := authHeader[0]

	if authorization == "" {
		log.Println("Token required")
		helpers.SendResponse(c, http.StatusBadRequest, "Token required", nil)

		return
	}

	jwtToken := strings.Split(authorization, " ")[1]
	userIdFromJwtToken, err := middlewares.ExtractJWTToken(jwtToken)
	userId := int(userIdFromJwtToken.(float64))

	if err != nil {
		log.Println("JWT error: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, err.Error(), nil)

		return
	}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	photoFromDatabase, err := photoQuery.Get(uint(userId))

	if err != nil {
		log.Println("Error not found: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, err.Error(), nil)

		return
	}

	if userId != int(photoFromDatabase.UserID) || photoId != int(photoFromDatabase.ID) {
		log.Println("Unauthorized")
		helpers.SendResponse(c, http.StatusUnauthorized, "Unauthorized", nil)

		return
	}

	var photo models.Photo

	if err := c.Bind(&photo); err != nil {
		log.Println("Binding error: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	file, err := c.FormFile("file")

	if err != nil {
		log.Println("Binding error: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	filePath, err := helpers.SaveFileToDir(c, file)

	if err != nil {
		log.Println("Save error: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	if err := helpers.RemoveFileFromDir(photoFromDatabase.PhotoURL); err != nil {
		log.Println("Error while deleting photo: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	photo.ID = uint(photoId)
	photo.PhotoURL = filePath

	if err := photoQuery.Update(&photo); err != nil {
		log.Println("failed to update user's photo: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	data := response.Photo{
		ID:       photo.ID,
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoURL: photo.PhotoURL,
		UserID:   uint(userId),
	}
	helpers.SendResponse(c, http.StatusOK, "Successfully update photo", data)
}

func (p *PhotoController) Delete(c *gin.Context) {
	authHeader := c.Request.Header["Authorization"]

	if authHeader == nil {
		log.Println("Header required")
		helpers.SendResponse(c, http.StatusBadRequest, "Header required", nil)

		return
	}

	authorization := authHeader[0]

	if authorization == "" {
		log.Println("Token required")
		helpers.SendResponse(c, http.StatusBadRequest, "Token required", nil)

		return
	}

	jwtToken := strings.Split(authorization, " ")[1]
	userIdFromJwtToken, err := middlewares.ExtractJWTToken(jwtToken)
	userId := int(userIdFromJwtToken.(float64))

	if err != nil {
		log.Println("JWT error: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, err.Error(), nil)

		return
	}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	photoFromDatabase, err := photoQuery.Get(uint(userId))

	if err != nil {
		log.Println("Error not found: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, err.Error(), nil)

		return
	}

	if userId != int(photoFromDatabase.UserID) || photoId != int(photoFromDatabase.ID) {
		log.Println("Unauthorized")
		helpers.SendResponse(c, http.StatusUnauthorized, "Error while deleting photo", nil)

		return
	}

	if err := helpers.RemoveFileFromDir(photoFromDatabase.PhotoURL); err != nil {
		log.Println("Error while deleting photo: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, err.Error(), nil)

		return
	}

	if err := photoQuery.Delete(uint(photoId)); err != nil {
		log.Println("Error while deleting photo")
		helpers.SendResponse(c, http.StatusUnauthorized, "Error while deleting photo", nil)

		return
	}

	helpers.SendResponse(c, http.StatusOK, "Successfully delete photo", nil)
}
