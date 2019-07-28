package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	roleService "reading-club-backend/service/role"
	userService "reading-club-backend/service/user"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// Login request body of login request
type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// User login user request
type User struct {
	UserName string
}

var UserRole string

// AuthMiddleWare jwt auth middle ware
func AuthMiddleWare() (*jwt.GinJWTMiddleware, error) {
	var port = os.Getenv("PORT")
	var identityKey = "name"
	if port == "" {
		port = "80"
	}

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     10,
		MaxRefresh:  30,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			userName := claims[identityKey].(string)
			return &User{
				UserName: userName,
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals Login
			if err := c.ShouldBindJSON(&loginVals); err != nil {
				fmt.Println(err)
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password

			UserRole = roleService.GetUserRole(loginVals.Username)
			if userService.ValidateUser(userID, password) == true {
				return &User{
					UserName: userID,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*User); ok {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		LoginResponse: func(c *gin.Context, code int, token string, t time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"code":   http.StatusOK,
				"token":  token,
				"expire": t.Format(time.RFC3339),
				"role":   UserRole,
			})
		},

		RefreshResponse: func(c *gin.Context, code int, token string, t time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"code":   http.StatusOK,
				"token":  token,
				"expire": t.Format(time.RFC3339),
			})
		},

		TokenLookup: "header: Authorization, query: token, cookie: jwt",

		TokenHeadName: "Bearer",

		TimeFunc: time.Now,

		SendCookie: true,
	})

	return authMiddleware, err
}
