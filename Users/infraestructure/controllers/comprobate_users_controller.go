package controllers

import (
	"net/http"

	"example.com/m/Users/application"
	"github.com/gin-gonic/gin"
)

type ComprobateUsersControllers struct {
	Services application.ComprobateUsers
}

func NewComprobateUsersControllers(uc application.ComprobateUsers)*ComprobateUsersControllers{
	return&ComprobateUsersControllers{Services: uc}
}

func(controller *ComprobateUsersControllers)Execute(c *gin.Context){
	result, err := controller.Services.ComprobateUsersProcess()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al obtener cantidad de usuarios",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Message": "El sistema cuenta con la sig. cantidad de usuarios",
		"Users": result,
	})
}