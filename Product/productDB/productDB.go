package productDB

import (
	"Product/models"
	"fmt"

	"gorm.io/gorm"
)

type ProductDB struct {
	Client interface{}
}

func (c *ProductDB) Get(id string) (*models.Products, error) {
	contact := &models.Products{}
	result := c.Client.(*gorm.DB).First(contact, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return contact, nil
}

//create method
//func (p *ProductDB) Create(c *models.Products) (*[]models.Products, error) {
//	Products = append(Products, *c)
//	return &Products, nil
//}
func (c *ProductDB) Create(contact *models.Products) (interface{}, error) {
	c.Client.(*gorm.DB).AutoMigrate(&models.Products{})
	result := c.Client.(*gorm.DB).Create(contact)
	if result.Error != nil {
		fmt.Println("------------->", result.Error)
		return nil, result.Error
	}
	return contact.ID, nil
}

//....
/*func (p *ProductDB) GetByName(s string) (*[]models.Products, error) {
	var Products1 []models.Products
	for index, item := range Products {
		if item.Name == s {
			Products1 = append(Products1, Products[index])
		}
	}
	return &Products1, nil
}*/

//Update Method
/*func (p *ProductDB) Update(s string, t *models.Products) (*[]models.Products, error) {
	for index, item := range Products {
		if item.Name == s {
			Products[index].Number = t.Number
			Products[index].Category = t.Category
			Products[index].Description = t.Description
		}
	}
	return &Products, nil
}*/

func (p *ProductDB) Update(s string, t *models.Products) (interface{}, error) {
	product := &models.Products{}
	p.Client.(*gorm.DB).Model(&product).Where("id=?", s).Updates(&models.Products{Name: t.Name, Number: t.Number, Category: t.Category, Description: t.Description})

	return product.Name, nil
}

func (c *ProductDB) Delete(id string) (interface{}, error) {
	result := c.Client.(*gorm.DB).Delete(&models.Products{}, id)
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
