package helpers

import "github.com/gin-gonic/gin"

type response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, status int, message string, data interface{}) {
	response := response{Message: message, Data: data}
	c.JSON(status, response)
}
