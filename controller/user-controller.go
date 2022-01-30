package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fahmialfareza/go_gonic_api/dto"
	"github.com/fahmialfareza/go_gonic_api/helper"
	"github.com/fahmialfareza/go_gonic_api/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userController) Update(context *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	errDTO := context.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	userID, ok := context.Get("user_id")
	if !ok {
		panic("invalid credential")
	}
	id, err := strconv.ParseUint(fmt.Sprintf("%v", userID), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	userUpdateDTO.ID = id
	user := c.userService.Update(userUpdateDTO)
	res := helper.BuildResponse(true, "OK!", user)
	context.JSON(http.StatusOK, res)
}

func (c *userController) Profile(context *gin.Context) {
	userID, ok := context.Get("user_id")
	if !ok {
		panic("invalid credential")
	}
	user := c.userService.Profile(fmt.Sprintf("%v", userID))
	res := helper.BuildResponse(true, "OK!", user)
	context.JSON(http.StatusOK, res)
}
