package controllers

import (
	"net/http"
	"strconv"

	"example.com/m/Users/application"
	"github.com/gin-gonic/gin"
)

type GetUserController struct {
	Services application.GetUser
}

func NewGetUserController(uc application.GetUser)*GetUserController{
	return &GetUserController{Services: uc}
}

func(controller GetUserController)Execute(c *gin.Context){
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

	results, err := controller.Services.GetUserProcess(id_number_32)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "No se pudo obtener el usuario",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Results": results,
	})
}