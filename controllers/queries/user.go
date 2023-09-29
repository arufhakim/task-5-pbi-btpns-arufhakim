package queries

import (
	"log"
	"task-5-pbi-btpns-arufhakim/database"
	"task-5-pbi-btpns-arufhakim/models"

	"gorm.io/gorm"
)

type UserQuery struct{}

var photo PhotoQuery
var user UserQuery

func (u *UserQuery) Save(user *models.User) error {
	return database.DB.Save(user).Error
}

func (u *UserQuery) Get(email string) (models.User, error) {
	var err error
	var user models.User

	if err = database.DB.Table("users").Where("email = ?", email).First(&user).Error; err != nil && err != gorm.ErrRecordNotFound {
		return user, err
	}

	if err == gorm.ErrRecordNotFound {
		return user, nil
	}

	return user, nil
}

func (u *UserQuery) Update(user *models.User) error {
	if err := database.DB.Model(user).Where("id = ?", user.ID).Updates(models.User{Username: user.Username, Email: user.Email}).Error; err != nil {
		return err
	}

	return nil
}

func (u *UserQuery) BeforeDelete(userId uint) error {
	photoId, err := photo.GetPhotoId(userId)

	if err != nil {
		log.Println("failed to get user's photo: ", err)
		return err
	}

	if err := photo.Delete(photoId); err != nil {
		log.Println("failed to delete user's photo: ", err)
		return err
	}

	return nil
}

func (u *UserQuery) Delete(userId uint) error {
	if err := user.BeforeDelete(userId); err != nil {
		return err
	}

	if err := database.DB.Unscoped().Delete(&models.User{}, userId).Error; err != nil {
		log.Println("failed to delete user", err)
		return err
	}

	return nil
}
