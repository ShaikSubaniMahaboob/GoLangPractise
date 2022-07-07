package main

import (
	"Product/handler"
	"Product/productDB"

	"github.com/gin-gonic/gin"
)

func main() {
	pdb := &productDB.ProductDB{}
	ph := &handler.ProductHandler{IProducts: pdb}
	r := gin.Default()
	r.GET("/Product/get", ph.GetHand())
	r.POST("/Product/Create", ph.CreateHand())
	r.GET("/Product/GetByName/:name", ph.GetByNameHand())
	r.PUT("/Product/Update/:name", ph.UpdateHand())
	r.DELETE("/Product/Delete/:name", ph.DeleteHand())
	r.Run()
}
