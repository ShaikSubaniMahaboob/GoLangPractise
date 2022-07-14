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

type ProductHandler struct {
	IProducts Interfaces.IProducts
	//m         *messaging.ProduceMessage
}

/*func (p *ProductHandler) GetHand() func(c *gin.Context) {
	return func(c *gin.Context) {
		product, _ := p.IProducts.Get(id)
		c.JSON(http.StatusPartialContent, *product)
	}
}*/

func (ch *ProductHandler) GetHand() func(*gin.Context) {
	return func(c *gin.Context) {
		if ch == nil || ch.IProducts == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in the handler",
			})
			glog.Errorln("ProductHandler or IProducts is nil")
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
		contact, err := ch.IProducts.Get(id)
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
		glog.Info("Product successfully fetched:", *contact)
		c.Abort()
	}
}

/*func (p *ProductHandler) CreateHand() func(c *gin.Context) {
	return func(c *gin.Context) {
		product := &models.Products{}
		//a:=productDB.Products

		err := json.NewDecoder(c.Request.Body).Decode(product)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": err,
			})
			c.Abort()
			return
		}
		id, _ := p.IProducts.Create(product)
		c.JSON(http.StatusCreated, gin.H{
			"status":  "success",
			"message": id,
		})
		c.Abort()
	}

}*/
func (ch *ProductHandler) CreateHand() func(*gin.Context) {
	return func(c *gin.Context) {
		if ch == nil || ch.IProducts == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in the handler",
			})
			glog.Errorln("ProductHandler or IProducts is nil")
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

		contact := &models.Products{}
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
		id, err := ch.IProducts.Create(contact)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error to store in database",
			})
			glog.Errorln(err)
			c.Abort()
			return
		}

		//produce
		//ch.m = &messaging.ProduceMessage{Brokers: "localhost:29092", Topic: "CONTACT_CREATION", Key: []byte(string(contact.ID)), Data: buf}
		//err = ch.m.Produce(context.Background())
		glog.Errorln(err)
		c.JSON(http.StatusCreated, gin.H{
			"status":  "success",
			"message": id,
		})
		glog.Info("Product successfully created:", id)
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
func (p *ProductHandler) UpdateHand() func(*gin.Context) {
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
		name := c.Params.ByName("id")
		result, _ := p.IProducts.Update(name, product)
		c.JSON(http.StatusOK, result)

	}
}

/*func (p *ProductHandler) DeleteHand() func(*gin.Context) {
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
		result, _ := p.IProducts.Delete(name)
		abc := productDB.Products

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": fmt.Sprint(result, " record deleted"),
			"Rem":     abc,
		})
		c.Abort()
	}
}*/

func (ch *ProductHandler) DeleteHand() func(*gin.Context) {
	return func(c *gin.Context) {
		if ch == nil || ch.IProducts == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in the handler",
			})
			glog.Errorln("ProductHandler or IProducts is nil")
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

		/*if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "id cannot be empty",
			})
			glog.Errorln("id cannot be empty")
			c.Abort()
			return
		}*/
		result, err := ch.IProducts.Delete(id)
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
		glog.Info("Product successfully deleted:", result)
		c.Abort()
	}
}
