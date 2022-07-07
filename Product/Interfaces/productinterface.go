package Interfaces

import "Product/models"

type IProducts interface {
	Get() (*models.Products, error)
	//Create() (*models.Products, error)
}
