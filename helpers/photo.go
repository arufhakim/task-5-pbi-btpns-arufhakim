package helpers

import (
	"fmt"
	"mime/multipart"
	"os"

	"github.com/gin-gonic/gin"
)

func SaveFileToDir(c *gin.Context, file *multipart.FileHeader) (string,error) {
	workingDir, err := os.Getwd()
	
	if err != nil {
		return "", err
	}

	filePath := fmt.Sprintf("%s/user-img/%s", workingDir, file.Filename) 
	c.SaveUploadedFile(file, filePath)

	return filePath,nil
}

func RemoveFileFromDir(filePath string) error {
	if err := os.Remove(filePath); err != nil {
		return err
	}

	return nil
}