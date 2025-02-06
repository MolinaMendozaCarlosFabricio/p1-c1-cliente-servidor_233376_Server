package routes

import (
	aplication "example.com/m/Products/application"
	"example.com/m/Products/infraestructure"
	"example.com/m/Products/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	ps := infraestructure.NewCreateProductRepoMySQL()
	cp_aplication := aplication.NewCreateProduct(ps)
	cp_controller := controllers.NewCreateProductController(*cp_aplication)

	gap_aplication := aplication.NewGetAllProducts(ps)
	gap_controller := controllers.NewGetAllProductsController(*gap_aplication)

	ep_aplication := aplication.NewEditProduct(ps)
	ep_controller := controllers.NewEditProductController(ep_aplication)

	dp_aplication := aplication.NewDeleteProduct(ps)
	dp_controller := controllers.NewDeleteProductController(dp_aplication)

	up_aplication := aplication.NewUpdatePrice(ps)
	up_controller := controllers.NewUpdatePriceController(*up_aplication)

	gldp_aplication := aplication.NewGetLastAddedProduct(ps)
	gldp_controller := controllers.NewGetLastAddedProductController(*gldp_aplication)
	products := r.Group("/products")
	{
		products.POST("/", cp_controller.Execute)
		products.GET("/", gap_controller.Execute)
		products.GET("/get_last_added_product", gldp_controller.Execute)
		products.PUT("/:id", ep_controller.Execute)
		products.PUT("/update_price/:id", up_controller.Execute)
		products.DELETE("/:id", dp_controller.Execute)
	}
}