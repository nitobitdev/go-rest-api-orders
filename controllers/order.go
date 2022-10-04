package controllers

import (
	"fmt"
	"go-rest-api-orders/models"
	"go-rest-api-orders/views"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderController struct {
	db *gorm.DB
}

func NewOrderController(db *gorm.DB) *OrderController {
	return &OrderController{
		db: db,
	}
}

func (o *OrderController) GetOrders(c *gin.Context) {
	orders := []models.Order{}

	if err := o.db.Find(&orders).Error; err != nil {
		c.JSON(http.StatusBadRequest, &views.ResponseError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if len(orders) <= 0 {
		c.JSON(http.StatusNotFound, &views.ResponseError{
			Status:  http.StatusNotFound,
			Message: "Data not found",
		})
		return
	}

	c.JSON(http.StatusOK, &views.ResponseSuccess{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    orders,
	})
	return
}

func (o *OrderController) CreateOrder(c *gin.Context) {
	order := models.Order{}
	item := models.Item{}
	arrayItem := []models.Item{}
	orderRequest := models.RequestOrder{}

	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		c.JSON(http.StatusBadRequest, &views.ResponseError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	order.Customer_name = orderRequest.Customer_name
	order.Ordered_at = orderRequest.Ordered_at

	if err := o.db.Create(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, &views.ResponseError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	for _, itemRequest := range orderRequest.Item {

		err := o.db.Last(&item).Error
		dataItem := o.db.Find(&arrayItem).Error
		if err != nil && dataItem != nil {
			c.JSON(http.StatusBadRequest, &views.ResponseError{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}

		if len(arrayItem) == 0 {
			item.Item_id = 1
		} else {
			item.Item_id = item.Item_id + 1
		}

		item.Order_id = order.Order_id
		item.Item_code = itemRequest.Item_code
		item.Description = itemRequest.Description
		item.Quantity = itemRequest.Quantity
		if err := o.db.Create(&item).Error; err != nil {
			c.JSON(http.StatusBadRequest, &views.ResponseError{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, &views.ResponseSuccess{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    orderRequest,
	})
	return
}

func (o *OrderController) UpdateOrder(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	order := models.Order{}
	orderRequest := models.RequestOrder{}
	item := []models.Item{}
	itemNew := models.Item{}
	arrayItem := []models.Item{}

	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		c.JSON(http.StatusBadRequest, &views.ResponseError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	err := o.db.Where("order_id = ?", id).First(&order).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, &views.ResponseError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	fmt.Println(order)
	order.Ordered_at = orderRequest.Ordered_at
	order.Customer_name = orderRequest.Customer_name

	if err := o.db.Save(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, &views.ResponseError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if err := o.db.Where("order_id = ?", id).Find(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, &views.ResponseError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if err := o.db.Delete(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, &views.ResponseError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	for _, itemRequest := range orderRequest.Item {
		err := o.db.Last(&itemNew).Error
		dataItem := o.db.Find(&arrayItem).Error
		if err != nil && dataItem != nil {
			c.JSON(http.StatusBadRequest, &views.ResponseError{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}

		if len(arrayItem) == 0 {
			itemNew.Item_id = 1
		} else {
			itemNew.Item_id = itemNew.Item_id + 1
		}

		itemNew.Order_id = order.Order_id
		itemNew.Item_code = itemRequest.Item_code
		itemNew.Description = itemRequest.Description
		itemNew.Quantity = itemRequest.Quantity
		if err := o.db.Create(&itemNew).Error; err != nil {
			c.JSON(http.StatusBadRequest, &views.ResponseError{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, &views.ResponseSuccess{
		Status:  http.StatusOK,
		Message: "Success",
	})
	return
}

func (o *OrderController) DeleteOrder(c *gin.Context) {
	order := models.Order{}
	item := []models.Item{}

	orderId, _ := strconv.Atoi(c.Param("id"))
	if err := o.db.Where("order_id = ?", orderId).Find(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, &views.ResponseError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if err := o.db.Where("order_id = ?", orderId).Find(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, &views.ResponseError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if err := o.db.Delete(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, &views.ResponseError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if err := o.db.Delete(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, &views.ResponseError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &views.ResponseSuccess{
		Status:  http.StatusOK,
		Message: "Success",
	})
}
