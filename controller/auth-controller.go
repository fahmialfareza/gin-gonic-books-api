package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthController interface is a contract what this controller can do
type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	// input the service that you need
	// this is where you put your service
}

// NewAuthController is a constructor for AuthController
func NewAuthController() AuthController {
	return &authController{}
}

// Login is a function that handles login
func (c *authController) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello Login",
	})
}

// Register is a function that handles register
func (c *authController) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello Register",
	})
}
