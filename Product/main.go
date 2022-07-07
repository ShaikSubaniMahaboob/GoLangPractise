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
	r.GET("/ping", ph.GetHand())
	r.Run()
}
