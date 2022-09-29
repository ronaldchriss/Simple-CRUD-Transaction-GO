package handler

import (
	"net/http"
	"test_aac/auth"
	"test_aac/helper"
	"test_aac/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	AuthService auth.Service
}

func NewUserHandler(userService user.Service, AuthService auth.Service) *userHandler {
	return &userHandler{userService, AuthService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMassage := gin.H{"errors": errors}

		response := helper.JsonResponse("Register Account Failed", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.JsonResponse("Register Account Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.AuthService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.JsonResponse("Register Account Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, token)
	response := helper.JsonResponse("Account Has Been Register", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMassage := gin.H{"errors": errors}

		response := helper.JsonResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMassage := gin.H{"errors": err.Error()}

		response := helper.JsonResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.AuthService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := helper.JsonResponse("Login Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, token)
	response := helper.JsonResponse("Successfuly Login", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckUsername(c *gin.Context) {
	var input user.CheckUsername

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMassage := gin.H{"errors": errors}

		response := helper.JsonResponse("Username Has Been Registered", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	UsernameAvail, err := h.userService.CheckUsername(input)
	if err != nil {
		errorMassage := gin.H{"errors": "Server Error"}

		response := helper.JsonResponse("Username Check Failed", http.StatusUnprocessableEntity, "error", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": UsernameAvail,
	}

	metamessage := "Username Has Been Registered"

	if UsernameAvail {
		metamessage = "Username Available"
	}

	response := helper.JsonResponse(metamessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) FetchUser(c *gin.Context) {
	User := c.MustGet("codeUser").(user.User)

	format := user.FormatUser(User, "")

	response := helper.JsonResponse("Success to Fetch Data User", http.StatusOK, "success", format)

	c.JSON(http.StatusOK, response)

}
