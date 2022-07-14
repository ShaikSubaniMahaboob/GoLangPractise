package main

import (
	"Product/handler"
	"Product/productDB"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

var (
	DBCONNECTION string = "host=localhost user=postgres password=subani@123 dbname=subani port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	PORT         string = ":50080"
)

func main() {
	if os.Getenv("PORT") != "" {
		PORT = os.Getenv("PORT")
	}

	if os.Getenv("DBCONNECTION") != "" {
		DBCONNECTION = os.Getenv("DBCONNECTION")
	}

	db, err := productDB.GetConnection(DBCONNECTION)

	if err != nil {
		glog.Fatal("Database Error:", err)
	}

	//cdb := &productDB.ProductDB{Client: db}
	//productHandler := &handler.ProductHandler{IProducts: cdb}

	//config.Connect()
	pdb := &productDB.ProductDB{Client: db}
	ph := &handler.ProductHandler{IProducts: pdb}

	cpd := &productDB.CustomersDB{Client: db}
	ch := &handler.CustomersHandler{ICustomers: cpd}

	opd := &productDB.OrdersDB{Client: db}
	oh := &handler.OrdersHandler{IOrders: opd}
	r := gin.Default()
	//product
	r.GET("/Product/get/:id", ph.GetHand())
	r.POST("/Product/Create", ph.CreateHand())

	r.PUT("/Product/Update/:id", ph.UpdateHand())
	r.DELETE("/Product/Delete/:id", ph.DeleteHand())
	//customer
	r.GET("/Customers/get/:id", ch.CustomerGetHand())
	r.POST("/Customers/Create", ch.CustomerCreateHand())
	r.PUT("/Customers/Update/:id", ch.UpdateCustomerHand())
	r.DELETE("/Customers/Delete/:id", ch.CustomerDeleteHand())
	//order
	r.GET("/Orders/get/:id", oh.OrderGetHand())
	r.POST("/Orders/Create", oh.OrderCreateHand())
	r.PUT("/Orders/Update/:orderid", oh.UpdateOrderHand())
	r.DELETE("/Orders/Delete/:id", oh.OrderDeleteHand())
	r.Run()

}
