package Interfaces

import "Product/models"

type ICustomers interface {
	Get() (*[]models.Customers, error)
	Create(*models.Customers) (*[]models.Customers, error)
	Delete(string) (string, error)
	GetById(string) (*[]models.Customers, error)
	Update(string, *models.Customers) (*[]models.Customers, error)
}
