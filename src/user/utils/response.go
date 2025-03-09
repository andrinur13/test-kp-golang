package utils

import (
	"net/http"

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
		Message: message,
		Data:    data,
	}

	c.JSON(http.StatusOK, response)
}
