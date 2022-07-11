package handler

import (
	"Product/Interfaces"
	"Product/models"
	"Product/productDB"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrdersHandler struct {
	IOrders Interfaces.IOrders
}

func (p *OrdersHandler) GetOrderHand() func(*gin.Context) {
	return func(c *gin.Context) {
		order, _ := p.IOrders.Get()
		c.JSON(http.StatusOK, order)
	}
}

func (p *OrdersHandler) CreateOrderHand() func(*gin.Context) {
	return func(c *gin.Context) {
		//_, err := ioutil.ReadAll(c.Request.Body)

		order := &models.Orders{}
		err := json.NewDecoder(c.Request.Body).Decode(order)
		//err = json.Unmarshal(buf, err)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": err,
			})
			c.Abort()
			return
		}

		id, _ := p.IOrders.Create(order)

		c.JSON(http.StatusCreated, gin.H{
			"status":  "success",
			"message": id,
		})
		c.Abort()

	}
}

func (p *OrdersHandler) DeleteOrderHand() func(*gin.Context) {
	return func(c *gin.Context) {
		orderid := c.Params.ByName("orderid")

		if orderid == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "id cannot be empty",
			})
			c.Abort()
			return
		}
		result, _ := p.IOrders.Delete(orderid)
		abc := productDB.Orders

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": fmt.Sprint(result, " record deleted"),
			"Rem":     abc,
		})
		c.Abort()
	}
}

func (p *OrdersHandler) GetByIdHand() func(*gin.Context) {
	return func(c *gin.Context) {

		orderid := c.Params.ByName("orderid")

		result, _ := p.IOrders.GetById(orderid)

		c.JSON(http.StatusOK, result)

	}
}

func (p *OrdersHandler) UpdateOrderHand() func(*gin.Context) {
	return func(c *gin.Context) {
		order := &models.Orders{}
		err := json.NewDecoder(c.Request.Body).Decode(order)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": err,
			})
			c.Abort()
			return
		}
		orderid := c.Params.ByName("orderid")
		result, _ := p.IOrders.Update(orderid, order)
		c.JSON(http.StatusOK, result)

	}
}
