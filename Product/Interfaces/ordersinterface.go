package Interfaces

import "Product/models"

type IOrders interface {
	Get() (*[]models.Orders, error)
	Create(*models.Orders) (*[]models.Orders, error)
	Delete(string) (string, error)
	GetById(string) (*[]models.Orders, error)
	Update(string, *models.Orders) (*[]models.Orders, error)
}
