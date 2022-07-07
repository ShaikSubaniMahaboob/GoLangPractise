package handler

import (
	"Product/Interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	IProducts Interfaces.IProducts
}

func (p *ProductHandler) GetHand() func(c *gin.Context) {
	return func(c *gin.Context) {
		product, _ := p.IProducts.Get()
		c.JSON(http.StatusPartialContent, *product)
	}
}

//func (p1 *ProductHandler) Create(Product *models.Products) (interface{}, error) {

//p1.IProducts.(*gorm.DB).AutoMigrate(&models.Products{})
//result := p1.IProducts.(*gorm.DB).Create(Product)
//if result.Error != nil {
//	fmt.Println("------------->", result.Error)
//	return nil, result.Error
//}
//return Product.Name, nil
/*router := gin.Default()

router.POST("/post", func(c *gin.Context) {

	ids := c.QueryMap("ids")
	names := c.PostFormMap("names")

	fmt.Printf("ids: %v; names: %v", ids, names)
})
router.Run(":8080")*/
//}
