package main

import (
	"fmt"
	// "github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=gorm dbname=adventuretime sslmode=disable password=")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&User{})
	fmt.Println("automigration complete")
}

type User struct {
	gorm.Model
	UserName   string
	CurrentHp int
	Xp         int
}

