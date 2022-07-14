package models

type Customers struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Number  string `json:"number"`
	Email   string `json:"email"`
	Address string `json:"address"`
}
