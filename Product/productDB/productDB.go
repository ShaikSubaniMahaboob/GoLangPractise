package productDB

import (
	"Product/models"
	"fmt"
)

type ProductDB struct {
}

var Products []models.Products

//get method
func (p *ProductDB) Get() (*[]models.Products, error) {

	//m := &models.Products{Name: "Apple", Number: "1890", Category: "MobilePhone", Description: "Apple2.0"}

	ErrorType := fmt.Errorf("there is Error in Products Database")
	return &Products, ErrorType
}

//create method
func (p *ProductDB) Create(c *models.Products) (*[]models.Products, error) {
	Products = append(Products, *c)
	return &Products, nil
}

//GetByName
func (p *ProductDB) GetByName(s string) (*[]models.Products, error) {
	var Products1 []models.Products
	for index, item := range Products {
		if item.Name == s {
			Products1 = append(Products1, Products[index])
		}
	}
	return &Products1, nil
}

//Update Method
func (p *ProductDB) Update(s string, t *models.Products) (*[]models.Products, error) {
	for index, item := range Products {
		if item.Name == s {
			Products[index].Number = t.Number
			Products[index].Category = t.Category
			Products[index].Description = t.Description
		}
	}
	return &Products, nil
}
func (p *ProductDB) Delete(s string) (string, error) {

	for index, item := range Products {
		if item.Name == s {
			Products = append(Products[:index], Products[index+1:]...)
		}
	}
	return s, nil
}
