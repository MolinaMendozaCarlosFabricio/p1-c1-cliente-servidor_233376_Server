package controllers

import (
	"net/http"
	"strconv"

	"example.com/m/Users/application"
	"github.com/gin-gonic/gin"
)

type DeleteUserController struct {
	Services application.DeleteUser
}

func NewDeleteUserController(uc application.DeleteUser)*DeleteUserController{
	return &DeleteUserController{Services: uc}
}

func(controller DeleteUserController)Execute(c *gin.Context){
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

	if err := controller.Services.DeleteUserProcess(id_number_32); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "No se pudo eliminar al usuario",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Message": "Usuario eliminado",
	})

}