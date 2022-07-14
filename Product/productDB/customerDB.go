package productDB

import (
	"Product/models"
	"fmt"

	"gorm.io/gorm"
)

//var (
//	ERROR_CONTACT_EXISTS = errors.New("customer already exists with the given email address")
//)

type CustomersDB struct {
	Client interface{}
}

func (c *CustomersDB) Get(id string) (*models.Customers, error) {
	contact := &models.Customers{}
	result := c.Client.(*gorm.DB).First(contact, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return contact, nil
}

func (c *CustomersDB) Create(contact *models.Customers) (interface{}, error) {
	c.Client.(*gorm.DB).AutoMigrate(&models.Customers{})
	result := c.Client.(*gorm.DB).Create(contact)
	if result.Error != nil {
		fmt.Println("------------->", result.Error)
		return nil, result.Error
	}
	return contact.ID, nil
}

//var Products []models.Products

/*func (p *CustomersDB) Update(s string, t *models.Customers) (interface{}, error) {
	for index, item := range Products {
		if item.Name == s {
			Products[index].Number = t.Number
			Products[index].Email = t.Email
			Products[index].Address = t.Address
		}
	}
	return &Products, nil
}*/
func (p *CustomersDB) Update(s string, t *models.Customers) (interface{}, error) {
	product := &models.Customers{}
	p.Client.(*gorm.DB).Model(&product).Where("id=?", s).Updates(&models.Customers{Name: t.Name, Number: t.Number, Email: t.Email, Address: t.Address})

	return product.Name, nil
}
func (c *CustomersDB) Delete(id string) (interface{}, error) {
	result := c.Client.(*gorm.DB).Delete(&models.Customers{}, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return result.RowsAffected, nil
}

/*func (p *ProductDB) Delete(s string) (string, error) {
for index, item := range Products {
		if item.Name == s {
			Products = append(Products[:index], Products[index+1:]...)
		}
	}
	return s, nil
}
*/
