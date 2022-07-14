package Interfaces

import "Product/models"

type ICustomers interface {
	Get(string) (*models.Customers, error)
	Create(*models.Customers) (interface{}, error)
	Update(string, *models.Customers) (interface{}, error)
	Delete(string) (interface{}, error)
}
