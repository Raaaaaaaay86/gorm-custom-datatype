package main

import (
	"encoding/json"
	"fmt"

	"github.com/raaaaaaaay86/gorm-custom-datatype/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := CreateMySQLGormConnection()

	// Create
	item := model.Item{
		Name:  "Book",
		Owner: []string{"Mary", "Jane"},
	}

	db.Create(&item)

	// Query
	queryItem := model.Item{}

	db.Where("id = ?", item.ID).First(&queryItem)

	marshaled, _ := json.MarshalIndent(queryItem, "", "\t")
	fmt.Println(string(marshaled))
}

func CreateMySQLGormConnection() *gorm.DB {
	db, err := gorm.Open(
		mysql.New(mysql.Config{
			DSN:                       "root:changeme@tcp(localhost:3307)/property?charset=utf8&parseTime=true&loc=Local",
			DefaultStringSize:         256,
			DisableDatetimePrecision:  true,
			DontSupportRenameIndex:    true,
			DontSupportRenameColumn:   true,
			SkipInitializeWithVersion: false,
		}),
		&gorm.Config{},
	)
	if err != nil {
		panic(fmt.Errorf("gorm initialization: MySQL connection failed: %s", err))
	}

	// db.AutoMigrate(&model.Item{})

	return db
}
