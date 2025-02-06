package controllers

import (
	"net/http"
	"strconv"

	aplication "example.com/m/Products/application"
	"github.com/gin-gonic/gin"
)

type EditProductController struct {
	Services *aplication.EditProduct
}

func NewEditProductController(uc *aplication.EditProduct)*EditProductController{
	return &EditProductController{Services: uc}
}

func(controller *EditProductController)Execute(c *gin.Context){
	var input struct {
		Name string `json:"name"`
		Price float32 `json:"price"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Entrada de datos no válida",
		})
		return
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

	if err := controller.Services.EditProductProcess(input.Name, input.Price, id_number_32); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "No se pudo editar el producto",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Producto editado",
	})
}