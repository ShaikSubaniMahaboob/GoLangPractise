package Interfaces

import "Product/models"

type IProducts interface {
	Get(string) (*models.Products, error)

	Create(*models.Products) (interface{}, error)
	Update(string, *models.Products) (interface{}, error)
	Delete(string) (interface{}, error)
}
