package middleware

import (
	"github.com/gin-gonic/gin"
)

//AllowCORS add response header to allow cors http request from client
func AllowCORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

		c.Next()
	}
}
