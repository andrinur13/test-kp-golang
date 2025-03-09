package utils

import (
	"github.com/gin-gonic/gin"
)

type Reponse struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func FormatResponse(c *gin.Context, status string, code int, message string, data interface{}) {
	response := Reponse{
		Status:  status,
		Code:    code,
		Message: message,
		Data:    data,
	}

	c.JSON(code, response)
}

func FormatErrorResponse(c *gin.Context, status string, code int, message string, data interface{}) {
	response := Reponse{
		Status:  status,
		Code:    code,
		Message: message,
		Data:    data,
	}

	c.JSON(code, response)
}
