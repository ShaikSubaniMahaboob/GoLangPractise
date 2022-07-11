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

type ProductHandler struct {
	IProducts Interfaces.IProducts
}

func (p *ProductHandler) GetHand() func(c *gin.Context) {
	return func(c *gin.Context) {
		product, _ := p.IProducts.Get()
		c.JSON(http.StatusPartialContent, *product)
	}
}
func (p *ProductHandler) CreateHand() func(c *gin.Context) {
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

}

func (p *ProductHandler) GetByNameHand() func(*gin.Context) {
	return func(c *gin.Context) {

		name := c.Params.ByName("name")

		result, _ := p.IProducts.GetByName(name)

		c.JSON(http.StatusOK, result)

	}
}

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
		name := c.Params.ByName("name")
		result, _ := p.IProducts.Update(name, product)
		c.JSON(http.StatusOK, result)

	}
}
func (p *ProductHandler) DeleteHand() func(*gin.Context) {
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
}
