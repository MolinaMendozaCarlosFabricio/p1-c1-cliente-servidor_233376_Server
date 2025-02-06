package controllers

import (
	"net/http"
	"strconv"

	"example.com/m/Users/application"
	"github.com/gin-gonic/gin"
)

type EditUserController struct {
	Services application.EditUser
}

func NewEditUserController(uc application.EditUser)*EditUserController{
	return &EditUserController{Services: uc}
}

func(controller EditUserController)Execute(c *gin.Context){
	var input struct {
		Username string `json:"username"`
		Email string `json:"email"`
		Password string `json:"password"`
	}

	id, error_param := c.Params.Get("id")
	if !error_param {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Error al obtener parámetros",
		})
		return
	}

	id_number, error_strconv := strconv.Atoi(id)
	if error_strconv != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Parámetro no válido",
		})
		return
	}
	id_number_32 := int32(id_number)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Entrada de datos no válida",
		})
		return
	}

	if err := controller.Services.EditUserProcess(
		id_number_32, input.Username, input.Email, input.Password,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "No se pudo editar el usuario",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Message": "Usuario editado",
	})
}