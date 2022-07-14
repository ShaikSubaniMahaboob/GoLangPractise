package handler

import (
	"Product/Interfaces"
	"Product/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

type CustomersHandler struct {
	ICustomers Interfaces.ICustomers
}

func (ch *CustomersHandler) CustomerGetHand() func(*gin.Context) {
	return func(c *gin.Context) {
		if ch == nil || ch.ICustomers == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in the handler",
			})
			glog.Errorln("CustomersHandler or ICustomers is nil")
			c.Abort()
			return
		}

		id, ok := c.Params.Get("id")

		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "id parameter not found",
			})
			glog.Errorln("id parameter not found")
			c.Abort()
			return
		}

		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "error in id param",
			})
			glog.Errorln("id cannot be empty")
			c.Abort()
			return
		}
		contact, err := ch.ICustomers.Get(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "error in fetching contact",
			})
			glog.Errorln(err)
			c.Abort()
			return
		}
		if contact == nil {
			c.JSON(http.StatusNoContent, nil)
			glog.Info("No content")
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, *contact)
		glog.Info("Customers successfully fetched:", *contact)
		c.Abort()
	}
}

func (ch *CustomersHandler) CustomerCreateHand() func(*gin.Context) {
	return func(c *gin.Context) {
		if ch == nil || ch.ICustomers == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in the handler",
			})
			glog.Errorln("CustomerHandler or ICustomers is nil")
			c.Abort()
			return
		}

		buf, err := ioutil.ReadAll(c.Request.Body)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in the body",
			})
			glog.Errorln(err)
			c.Abort()
			return
		}

		contact := &models.Customers{}
		err = json.Unmarshal(buf, contact)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in body json format",
			})
			glog.Errorln(err)
			c.Abort()
			return
		}
		name, err := ch.ICustomers.Create(contact)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error to store in database",
			})
			glog.Errorln(err)
			c.Abort()
			return
		}

		glog.Errorln(err)
		c.JSON(http.StatusCreated, gin.H{
			"status":  "success",
			"message": name,
		})
		glog.Info("customers successfully created:", name)
		c.Abort()
	}
}

/*func (p *ProductHandler) GetByNameHand() func(*gin.Context) {
	return func(c *gin.Context) {

		name := c.Params.ByName("name")

		result, _ := p.IProducts.GetByName(name)

		c.JSON(http.StatusOK, result)

	}
}

/*func (p *ProductHandler) UpdateHand() func(*gin.Context) {
	return func(c *gin.Context) {
		product := &models.Products{}
		err := json.NewDecoder(c.Request.Body).Decode(product)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": err,
			})
			c.Abort()
			return
		}
		name := c.Params.ByName("name")
		result, _ := p.IProducts.Update(name, product)
		c.JSON(http.StatusOK, result)

	}
}*/

func (p *CustomersHandler) UpdateCustomerHand() func(*gin.Context) {
	return func(c *gin.Context) {
		product := &models.Customers{}
		err := json.NewDecoder(c.Request.Body).Decode(product)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": err,
			})
			c.Abort()
			return
		}
		name := c.Params.ByName("id")
		result, _ := p.ICustomers.Update(name, product)
		c.JSON(http.StatusOK, result)

	}
}

func (ch *CustomersHandler) CustomerDeleteHand() func(*gin.Context) {
	return func(c *gin.Context) {
		if ch == nil || ch.ICustomers == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in the handler",
			})
			glog.Errorln("CustomerHandler or ICustomers is nil")
			c.Abort()
			return
		}

		id, ok := c.Params.Get("id")

		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "id parameter not found",
			})
			glog.Errorln("id parameter not found")
			c.Abort()
			return
		}
		result, err := ch.ICustomers.Delete(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "error deleting contact",
			})
			glog.Errorln(err)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": fmt.Sprint(result, " record deleted"),
		})
		glog.Info("Customer successfully deleted:", result)
		c.Abort()
	}
}
