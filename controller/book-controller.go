package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fahmialfareza/go_gonic_api/dto"
	"github.com/fahmialfareza/go_gonic_api/entity"
	"github.com/fahmialfareza/go_gonic_api/helper"
	"github.com/fahmialfareza/go_gonic_api/service"
	"github.com/gin-gonic/gin"
)

type BookController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type bookController struct {
	bookService service.BookService
	jwtService  service.JWTService
}

func NewBookController(bookServ service.BookService, jwtServ service.JWTService) BookController {
	return &bookController{
		bookService: bookServ,
		jwtService:  jwtServ,
	}
}

func (c *bookController) All(context *gin.Context) {
	var books []entity.Book = c.bookService.All()
	res := helper.BuildResponse(true, "OK!", books)
	context.JSON(http.StatusOK, res)
}

func (c *bookController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var book entity.Book = c.bookService.FindByID(id)
	if (book == entity.Book{}) {
		res := helper.BuildErrorResponse("Book not found", "No data with given id", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	} else {
		res := helper.BuildResponse(true, "OK!", book)
		context.JSON(http.StatusOK, res)
	}
}

func (c *bookController) Insert(context *gin.Context) {
	var bookCreateDTO dto.BookCreateDTO
	errDTO := context.ShouldBind(&bookCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	} else {
		userID, ok := context.Get("user_id")
		if !ok {
			panic("invalid credential")
		}
		convertedUserID, err := strconv.ParseUint(fmt.Sprintf("v", userID), 10, 64)
		if err == nil {
			bookCreateDTO.UserID = convertedUserID
		}
		result := c.bookService.Insert(bookCreateDTO)
		response := helper.BuildResponse(true, "OK!", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *bookController) Update(context *gin.Context) {
	var bookUpdateDTO dto.BookUpdateDTO
	errDTO := context.ShouldBind(&bookUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	userID, ok := context.Get("user_id")
	if !ok {
		panic("invalid credential")
	}
	convertedUserID, err := strconv.ParseUint(fmt.Sprintf("v", userID), 10, 64)
	if err == nil {
		bookUpdateDTO.UserID = convertedUserID
	}
	if c.bookService.IsAllowedToEdit(string(convertedUserID), bookUpdateDTO.ID) {
		res := c.bookService.Update(bookUpdateDTO)
		response := helper.BuildResponse(true, "OK!", res)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("Failed to process request", "You are not allowed to edit this book", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusForbidden, response)
	}
}

func (c *bookController) Delete(context *gin.Context) {
	var book entity.Book
	id, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", "No param id were found", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	book.ID = id

	userID, ok := context.Get("user_id")
	if !ok {
		panic("invalid credential")
	}
	convertedUserID, err := strconv.ParseUint(fmt.Sprintf("v", userID), 10, 64)
	if err != nil {
		panic(err)
	}

	if c.bookService.IsAllowedToEdit(string(convertedUserID), book.ID) {
		c.bookService.Delete(book)
		response := helper.BuildResponse(true, "Deleted!", helper.EmptyObj{})
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("Failed to process request", "You are not allowed to delete this book", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusForbidden, response)
	}
}
