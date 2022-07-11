package productDB

import (
	"Product/models"
	"fmt"
)

type OrdersDB struct {
}

var Orders []models.Orders

func (p *OrdersDB) Get() (*[]models.Orders, error) {

	ErrorType := fmt.Errorf("there is Error in Products Database")
	return &Orders, ErrorType
}

func (p *OrdersDB) Create(c *models.Orders) (*[]models.Orders, error) {
	//pd := &models.Products{}

	Orders = append(Orders, *c)
	return &Orders, nil
}

func (p *OrdersDB) Delete(s string) (string, error) {

	for index, item := range Orders {
		if item.OrderID == s {
			Orders = append(Orders[:index], Orders[index+1:]...)
		}
	}
	return s, nil
}

func (p *OrdersDB) GetById(s string) (*[]models.Orders, error) {
	var Orders1 []models.Orders
	for index, item := range Orders {
		if item.OrderID == s {
			Orders1 = append(Orders1, Orders[index])
		}
	}
	return &Orders1, nil
}

func (p *OrdersDB) Update(s string, t *models.Orders) (*[]models.Orders, error) {
	for index, item := range Orders {
		if item.OrderID == s {
			Orders[index].ProductID = t.ProductID
			Orders[index].CustomerID = t.CustomerID
		}
	}
	return &Orders, nil
}
