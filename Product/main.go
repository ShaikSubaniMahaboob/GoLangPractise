package main

import (
	"Product/handler"
	"Product/productDB"

	"github.com/gin-gonic/gin"
)

func main() {
	//config.Connect()
	pdb := &productDB.ProductDB{}
	ph := &handler.ProductHandler{IProducts: pdb}
	cpd := &productDB.CustomersDB{}
	ch := &handler.CustomersHandler{ICustomer: cpd}
	opd := &productDB.OrdersDB{}
	oh := &handler.OrdersHandler{IOrders: opd}
	r := gin.Default()

	r.GET("/Product/get", ph.GetHand())
	r.POST("/Product/Create", ph.CreateHand())
	r.GET("/Product/GetByName/:name", ph.GetByNameHand())
	r.PUT("/Product/Update/:name", ph.UpdateHand())
	r.DELETE("/Product/Delete/:name", ph.DeleteHand())

	r.GET("/Customers/get", ch.GetCustomerHand())
	r.GET("/Customers/GetById/:name", ch.GetByIdHand())
	r.POST("/Customers/Create", ch.CreateCustomerHand())
	r.PUT("/Customers/Update/:name", ch.UpdateCustomerHand())
	r.DELETE("/Customers/Delete/:name", ch.DeleteCustomerHand())

	r.GET("/Orders/get", oh.GetOrderHand())
	r.GET("/Orders/GetById/:orderid", oh.GetByIdHand())
	r.POST("/Orders/Create", oh.CreateOrderHand())
	r.PUT("/Orders/Update/:orderid", oh.UpdateOrderHand())
	r.DELETE("/Orders/Delete/:orderid", oh.DeleteOrderHand())
	r.Run()
}
