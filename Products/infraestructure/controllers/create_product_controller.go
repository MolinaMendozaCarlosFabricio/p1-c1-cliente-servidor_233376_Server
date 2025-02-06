package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	aplication "example.com/m/Products/application"
)

type CreateProductController struct{
	useCase aplication.CreateProduct
}

func NewCreateProductController(uc aplication.CreateProduct) *CreateProductController{
	return &CreateProductController{useCase: uc}
}

func(controller CreateProductController) Execute(c *gin.Context){
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

	if err := controller.useCase.CreateProductProcess(input.Name, input.Price); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al registrar uxu",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Producto registrado",
	})
}