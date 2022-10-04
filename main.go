package main

import (
	"go-rest-api-orders/config"
	"go-rest-api-orders/controllers"
	"go-rest-api-orders/handlers"
)

func main() {

	db := config.ConnectGorm()
	orderController := controllers.NewOrderController(db)
	router := handlers.NewRouter(orderController)
	router.Start(":5000")

}
