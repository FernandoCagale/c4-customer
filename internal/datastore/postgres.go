package datastore

import (
	"fmt"
	"github.com/FernandoCagale/c4-customer/pkg/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"os"
)

func NewPostgres() (*gorm.DB, error) {
	fmt.Println(os.Getenv("POSTGRES_ADDRS"))
	db, err := gorm.Open("postgres", os.Getenv("POSTGRES_ADDRS"))
	if err != nil {
		return nil, err
	}

	db.LogMode(false)
	db.SingularTable(true)
	db.AutoMigrate(&entity.Customer{})

	return db, nil
}
