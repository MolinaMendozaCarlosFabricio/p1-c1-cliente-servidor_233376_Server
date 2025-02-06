package controllers

import (
	"net/http"
	"strconv"

	aplication "example.com/m/Products/application"
	"github.com/gin-gonic/gin"
)

type DeleteProductController struct{
	Services *aplication.DeleteProduct
}

func NewDeleteProductController(uc *aplication.DeleteProduct)*DeleteProductController{
	return &DeleteProductController{Services: uc}
}

func(controller DeleteProductController)Execute(c *gin.Context){
	str_id, bool_err := c.Params.Get("id")
	if !bool_err {
		c.JSON(404, gin.H{"Error":"falta el parámetro ID"})
		return
	}

	id, err_err := strconv.Atoi(str_id)
	if err_err != nil {
		c.JSON(502, gin.H{"Error":"Parámetro id no válido"})
		return
	}
	id_32 := int32(id)

	if err := controller.Services.DeleteProductProcess(id_32); err != nil {
		c.JSON(502, gin.H{"Error":"Error al ejecutar método"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Message": "Producto eliminado",
	})
}
func DeleteProduct(c *gin.Context){
	str_id, bool_err := c.Params.Get("id")
	if !bool_err {
		c.JSON(404, gin.H{"Error":"falta el parámetro ID"})
		return
	}

	_, err_err := strconv.ParseInt(str_id, 10, 32)
	if err_err != nil {
		c.JSON(502, gin.H{"Error":"Parámetro id no válido"})
		return
	}
	
	c.JSON(http.StatusAccepted, gin.H{
		"Message": "Producto eliminado",
	})
}