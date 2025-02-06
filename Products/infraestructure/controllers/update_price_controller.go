package controllers

import (
	"net/http"
	"strconv"

	aplication "example.com/m/Products/application"
	"github.com/gin-gonic/gin"
)

type UpdatePriceController struct {
	Service aplication.UpdatePrice
}

func NewUpdatePriceController(uc aplication.UpdatePrice)*UpdatePriceController{
	return&UpdatePriceController{Service: uc}
}

func(controller *UpdatePriceController)Execute(c *gin.Context){
	id, err := c.Params.Get("id")
	if !err {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Error al obtener parámetro ID",
		})
	}

	id_int, err_strconv := strconv.Atoi(id)

	if err_strconv != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "La id debe ser un número entero",
		})
	}

	var input struct {
		Price float32 `json:"price"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "El cuerpo de la consulta está mal",
		})
	}

	if err := controller.Service.UpdatePriceProcess(int32(id_int), input.Price); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al actualizar precio",
		})
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Message": "Precio actualizado para el producto",
	})

}