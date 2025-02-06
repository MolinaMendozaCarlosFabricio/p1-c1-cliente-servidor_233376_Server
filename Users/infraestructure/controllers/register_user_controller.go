package controllers

import (
	"net/http"

	"example.com/m/Users/application"
	"github.com/gin-gonic/gin"
)

type RegisterUserController struct {
	Services application.RegisterUser
}

func NewRegisterUserController(uc application.RegisterUser)*RegisterUserController{
	return &RegisterUserController{Services: uc}
}

func (controller *RegisterUserController)Execute(c *gin.Context){
	var input struct {
		Username string `json:"username"`
		Email string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Entrada de datos no válida",
		})
		return
	}
	if err := controller.Services.RegisterUserProcess(input.Username, input.Email, input.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al registrar producto",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario registrado",
	})
}