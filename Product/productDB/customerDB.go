package productDB

import (
	"Product/models"
	"fmt"
)

type CustomersDB struct {
}

var Customers []models.Customers

func (p *CustomersDB) Get() (*[]models.Customers, error) {

	ErrorType := fmt.Errorf("there is Error in Products Database")
	return &Customers, ErrorType
}

func (p *CustomersDB) Create(c *models.Customers) (*[]models.Customers, error) {
	//pd := &models.Products{}

	Customers = append(Customers, *c)
	return &Customers, nil
}

func (p *CustomersDB) Delete(s string) (string, error) {

	for index, item := range Customers {
		if item.Name == s {
			Customers = append(Customers[:index], Customers[index+1:]...)
		}
	}
	return s, nil
}

func (p *CustomersDB) GetById(s string) (*[]models.Customers, error) {
	var Customers1 []models.Customers
	for index, item := range Customers {
		if item.Name == s {
			Customers1 = append(Customers1, Customers[index])
		}
	}
	return &Customers1, nil
}

func (p *CustomersDB) Update(s string, t *models.Customers) (*[]models.Customers, error) {
	for index, item := range Customers {
		if item.Name == s {
			Customers[index].Number = t.Number
			Customers[index].Email = t.Email
			Customers[index].Address = t.Address

		}
	}
	return &Customers, nil
}
