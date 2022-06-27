package main

import "fmt"

func main() {
	p := &person{}
	p.Number = 100
	p.Name = "subani"
	p.Email = "subani@gmail.com"
	addr := &Address{Line1: "Flat No 202,Maharaj Apartments", City: "Bangalore", State: "Karnataka", Country: "India", PinCode: "560016"}
	p.Address = addr
	fmt.Println(p)
	fmt.Println(p.Address)
	p.LinkedIn = "subani"
	p.Twitter = "subani3454"
	p.InstaGram = "subaniInsta"
	p.FaceBook = "subani shaik"
	a := &Address{}
	a.Line1 = "Islampeta"
	a.City = "Ponnur"
	a.State = "Aadhrapradesh"
	a.Country = "India"
	a.PinCode = "522124"
	fmt.Println(a)

}

type person struct {
	Number  int
	Name    string
	Email   string
	Address *Address
	SocialMedia
}
type Address struct {
	Line1, City, State, Country, PinCode string
}
type SocialMedia struct {
	LinkedIn, Twitter, FaceBook, InstaGram string
}

func (p *person) SetPerson(number int, name, email string, address *Address) {
	p.Number = number
	p.Name = name
	p.Email = email
	p.Address = address

}

func (a *Address) SetAddess(line1, city, state, country, pincode string) {
	a.Line1 = line1
	a.City = city
	a.State = state
	a.Country = country
	a.PinCode = pincode
}
