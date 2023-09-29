package queries

import (
	"task-5-pbi-btpns-arufhakim/database"
	"task-5-pbi-btpns-arufhakim/models"
)

type PhotoQuery struct {}

func (p *PhotoQuery) Get(userId uint) (models.Photo, error) {
	var photo models.Photo

	if err := database.DB.Where("user_id = ?", userId).First(&photo).Error; err != nil {
		return photo, err
	}

	return photo, nil
}

func (p *PhotoQuery) GetPhotoId(userId uint) (uint, error) {
	var id uint
	if err := database.DB.Table("photos").Select("id").Where("user_id = ?", userId).Scan(&id).Error; err != nil {
		return 0, err
	}

	return id, nil
}

func (p *PhotoQuery) Delete(photoId uint) error {
	if err := database.DB.Unscoped().Delete(&models.Photo{}, photoId).Error; err != nil {
		return err
	}

	return nil
}

func (p *PhotoQuery) Save(photo *models.Photo) error {
	return database.DB.Save(photo).Error
}

func (p *PhotoQuery) Update(photo *models.Photo) error {
	if err := database.DB.Model(photo).Where("id = ?", photo.ID).Updates(models.Photo{
		Title: photo.Title,
		Caption: photo.Caption,
		PhotoURL: photo.PhotoURL,
	}).Error; err != nil {
		return err
	}

	return nil
}