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

type CustomersHandler struct {
	ICustomer Interfaces.ICustomers
}

func (p *CustomersHandler) GetCustomerHand() func(*gin.Context) {
	return func(c *gin.Context) {
		order, _ := p.ICustomer.Get()
		c.JSON(http.StatusOK, order)
	}
}

func (p *CustomersHandler) CreateCustomerHand() func(*gin.Context) {
	return func(c *gin.Context) {

		customer := &models.Customers{}
		err := json.NewDecoder(c.Request.Body).Decode(customer)
		//err = json.Unmarshal(buf, err)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": err,
			})
			c.Abort()
			return
		}

		id, _ := p.ICustomer.Create(customer)

		c.JSON(http.StatusCreated, gin.H{
			"status":  "success",
			"message": id,
		})
		c.Abort()

	}
}

func (p *CustomersHandler) DeleteCustomerHand() func(*gin.Context) {
	return func(c *gin.Context) {
		name := c.Params.ByName("name")

		if name == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "id cannot be empty",
			})
			c.Abort()
			return
		}
		result, _ := p.ICustomer.Delete(name)
		//abc := productDb.Customers
		abc := productDB.Customers
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": fmt.Sprint(result, " record deleted"),
			"Rem":     abc,
		})
		c.Abort()
	}
}

func (p *CustomersHandler) GetByIdHand() func(*gin.Context) {
	return func(c *gin.Context) {

		name := c.Params.ByName("name")

		result, _ := p.ICustomer.GetById(name)

		c.JSON(http.StatusOK, result)

	}
}

func (p *CustomersHandler) UpdateCustomerHand() func(*gin.Context) {
	return func(c *gin.Context) {
		customer := &models.Customers{}
		err := json.NewDecoder(c.Request.Body).Decode(customer)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": err,
			})
			c.Abort()
			return
		}
		name := c.Params.ByName("name")
		result, _ := p.ICustomer.Update(name, customer)
		c.JSON(http.StatusOK, result)

	}
}
