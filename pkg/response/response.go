package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorDetail struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
}

type Envelope struct {
	Success bool          `json:"success"`
	Data    interface{}   `json:"data,omitempty"`
	Error   string        `json:"error,omitempty"`
	Code    string        `json:"code,omitempty"`
	Details []ErrorDetail `json:"details,omitempty"`
}

func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Envelope{Success: true, Data: data})
}

func Fail(c *gin.Context, status int, msg string) {
	c.JSON(status, Envelope{Success: false, Error: msg})
}

func FailWithCode(c *gin.Context, status int, msg string, code string) {
	c.JSON(status, Envelope{Success: false, Error: msg, Code: code})
}

func FailWithDetail(c *gin.Context, status int, msg string, details []ErrorDetail) {
	c.JSON(status, Envelope{
		Success: false,
		Error:   msg,
		Details: details,
	})
}

func FailWithCodeAndDetail(c *gin.Context, status int, msg string, code string, details []ErrorDetail) {
	c.JSON(status, Envelope{
		Success: false,
		Error:   msg,
		Code:    code,
		Details: details,
	})
}
