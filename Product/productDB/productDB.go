package productDB

import (
	"Product/models"
	"fmt"
)

type ProductDB struct {
}

func (p *ProductDB) Get() (*models.Products, error) {

	m := &models.Products{Name: "Subani", Number: "9701304581", Category: "Iphone", Description: "AboutProject"}
	ErrorType := fmt.Errorf("Error in Product Database")
	return m, ErrorType
}
