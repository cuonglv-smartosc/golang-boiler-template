
package response

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type Envelope struct {
    Success bool        `json:"success"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
}

func OK(c *gin.Context, data interface{}) {
    c.JSON(http.StatusOK, Envelope{Success: true, Data: data})
}

func Fail(c *gin.Context, status int, msg string) {
    c.JSON(status, Envelope{Success: false, Error: msg})
}
