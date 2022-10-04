package handlers

import (
	"go-rest-api-orders/controllers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	order *controllers.OrderController
}

func NewRouter(order *controllers.OrderController) *Router {
	return &Router{order: order}
}

func (r *Router) Start(port string) {
	router := gin.Default()

	router.GET("/orders", r.order.GetOrders)
	router.POST("/orders", r.order.CreateOrder)
	router.PUT("/orders/:id", r.order.UpdateOrder)
	router.DELETE("/orders/:id", r.order.DeleteOrder)
	router.Run(port)
}
