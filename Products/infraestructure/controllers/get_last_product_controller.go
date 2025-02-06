package controllers

import (
	"net/http"
	"strconv"
	"time"

	aplication "example.com/m/Products/application"
	"github.com/gin-gonic/gin"
)

type GetLastAddedProductController struct {
	Service aplication.GetLastAddedProduct
}

func NewGetLastAddedProductController(uc aplication.GetLastAddedProduct)*GetLastAddedProductController{
	return&GetLastAddedProductController{Service: uc}
}

func(controller *GetLastAddedProductController)Execute(c *gin.Context){
	c.Header("Content-Type", "application/json")
	c.Header("Transfer-Encoding", "chunked")	

	timeout := time.After(30 * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	var lastProductID int64

	for {
		select {
		case <-timeout:
			c.JSON(http.StatusRequestTimeout, gin.H{
				"Error": "Timeout: No se encontraron actualizaciones",
			})
			return

		case <-ticker.C:
			result, err := controller.Service.GetLastAddedProductProcess()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"Error": "Error al obtener el último producto registrado",
				})
				return
			}

			if result.ID > 0 && result.ID != int32(lastProductID) {
				lastProductID = int64(result.ID)

				idString := strconv.Itoa(int(result.ID))
				priceString := strconv.FormatFloat(float64(result.Price), 'f', 2, 32)

				payload := map[string]string{
					"Message":      "Último producto agregado",
					"Product_id":   idString,
					"Product_name": result.Name,
					"Product_price": priceString,
				}

				c.JSON(http.StatusOK, payload)
				return
			}
		}
	}
}