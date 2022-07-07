package Interfaces

import "Product/models"

type IProducts interface {
	Get() (*[]models.Products, error)
	Create(*models.Products) (*[]models.Products, error)
	GetByName(string) (*[]models.Products, error)
	Update(string, *models.Products) (*[]models.Products, error)
	Delete(string) (string, error)
}
