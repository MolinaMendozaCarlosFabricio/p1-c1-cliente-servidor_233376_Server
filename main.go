package main

import (
	product_routes "example.com/m/Products/infraestructure/routes"
	user_routes"example.com/m/Users/infraestructure/routes"
	"github.com/gin-gonic/gin"
)

func main () {
	r := gin.Default()

	product_routes.ProductRoutes(r)
	user_routes.UserRoutes(r)

	r.Run()
}