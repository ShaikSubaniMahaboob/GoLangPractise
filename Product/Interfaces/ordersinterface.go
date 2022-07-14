package Interfaces

import "Product/models"

type IOrders interface {
	Get(string) (*models.Orders, error)
	Create(*models.Orders) (interface{}, error)
	Update(string, *models.Orders) (interface{}, error)
	Delete(string) (interface{}, error)
}
