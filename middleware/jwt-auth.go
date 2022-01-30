package middleware

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/fahmialfareza/go_gonic_api/helper"
	"github.com/fahmialfareza/go_gonic_api/service"
	"github.com/gin-gonic/gin"
)

// AuthorizeJWT is a middleware for JWT
func AuthorizeJWT(jwtServ service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response := helper.BuildErrorResponse("Failed to process request", "No token found", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		token, err := jwtServ.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claims["user_id"])
			log.Println("Claim[issuer]: ", claims["issuer"])
			c.Set("user_id", claims["user_id"])
		} else {
			log.Println(err)
			response := helper.BuildErrorResponse("Token is not valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
	}
}
