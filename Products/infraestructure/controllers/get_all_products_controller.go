package controllers

import (
	"net/http"

	aplication "example.com/m/Products/application"
	"github.com/gin-gonic/gin"
)

type GetAllProductsController struct {
	Services aplication.GetAllProducts
}

func NewGetAllProductsController(uc aplication.GetAllProducts)*GetAllProductsController{
	return &GetAllProductsController{Services: uc}
}

func(controller *GetAllProductsController)Execute(c *gin.Context){
	result, err := controller.Services.GetAllProductProcess()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al obtener productos",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Response": result,
	})
}