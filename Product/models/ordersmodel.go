package models

type Orders struct {
	OrderID    string `json:"orderid"`
	ProductID  string `json:"productid"`
	CustomerID string `json:"customerid"`
}
