package routers

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	group := r.Group("/auth")
	{
		group.GET("/health", func(c *gin.Context) { c.Status(200) })
	}
}
