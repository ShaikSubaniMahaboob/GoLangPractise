package productDB

import (
	"Product/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

var (
	ERROR_CONTACT_EXISTS = errors.New("order already exists with the given email address")
)

type OrdersDB struct {
	Client interface{}
}

func (c *OrdersDB) Get(id string) (*models.Orders, error) {
	contact := &models.Orders{}
	result := c.Client.(*gorm.DB).First(contact, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return contact, nil
}

func (c *OrdersDB) Create(contact *models.Orders) (interface{}, error) {
	c.Client.(*gorm.DB).AutoMigrate(&models.Orders{})
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
func (p *OrdersDB) Update(s string, t *models.Orders) (interface{}, error) {
	product := &models.Orders{}
	p.Client.(*gorm.DB).Model(&product).Where("ID=?", s).Updates(&models.Orders{ProductID: t.ProductID, CustomerID: t.CustomerID, OrderID: t.OrderID})

	return product.ID, nil
}
func (c *OrdersDB) Delete(orderid string) (interface{}, error) {
	result := c.Client.(*gorm.DB).Delete(&models.Orders{}, orderid)
	if result.Error != nil {
		return nil, result.Error
	}
	return result.RowsAffected, nil
}
