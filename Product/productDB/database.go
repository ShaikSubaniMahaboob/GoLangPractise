package productDB

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	MAXRETRIES = 10
	MAXWAIT    = time.Second * 2
)

//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
func GetConnection(url string) (interface{}, error) {
	count := 1
retry:
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		if count <= MAXRETRIES {
			count = count + 1
			time.Sleep(MAXWAIT)
			fmt.Printf("Trying to connect to the database %v number of times\n", count)
			goto retry
		}
		return db, err
	}
	return db, err
}

//model:= model{ID: "1" Name: "Apple", Number: "10011", Category:"MobileNumber" ,Description:"Apple1.1"}

//result := db.Create(&models)
