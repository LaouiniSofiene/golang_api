package controller

import (
	"net/http"
	"strconv"

	"github.com/LaouiniSofiene/golang_api/dto"
	"github.com/LaouiniSofiene/golang_api/entity"
	"github.com/LaouiniSofiene/golang_api/helper"
	"github.com/LaouiniSofiene/golang_api/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	GetAllUsers(c *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(bookServ service.UserService) UserController {
	return &userController{
		userService: bookServ,
	}
}

func (c *userController) All(context *gin.Context) {
	var users []entity.User = c.userService.All()
	res := helper.BuildResponse(true, "OK", users)
	context.JSON(http.StatusOK, res)
}

func (c *userController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
	}

	var user entity.User = c.userService.FindByID(id)
	if (user == entity.User{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", user)
		context.JSON(http.StatusOK, res)
	}
}

func (c *userController) Insert(context *gin.Context) {
	var userCreateDTO dto.UserCreateDTO
	errDTO := context.ShouldBind(&userCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		if !c.userService.IsDuplicateEmail(userCreateDTO.Email) {
			response := helper.BuildErrorResponse("Failed to process request", "Duplicated email", helper.EmptyObj{})
			context.JSON(http.StatusConflict, response)
		} else {
			result := c.userService.Insert(userCreateDTO)
			response := helper.BuildResponse(true, "OK", result)
			context.JSON(http.StatusCreated, response)
		}

	}
}

func (c *userController) Update(context *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	errDTO := context.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		// if !c.userService.IsDuplicateEmail(userUpdateDTO.Email) {
		// 	response := helper.BuildErrorResponse("Failed to process request", "Duplicated email", helper.EmptyObj{})
		// 	context.JSON(http.StatusConflict, response)
		// } else {
		result := c.userService.Update(userUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
		// }
	}
}

func (c *userController) Delete(context *gin.Context) {
	var user entity.User
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	user.ID = id
	c.userService.Delete(user)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)
}

func (c *userController) GetAllUsers(context *gin.Context) {
	pagination := helper.GeneratePaginationFromRequest(context)
	var user entity.User
	userLists, err := c.userService.GetAllUsers(&user, &pagination)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}
	context.JSON(http.StatusOK, gin.H{
		"data": userLists,
	})
}
