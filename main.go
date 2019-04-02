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
	db.AutoMigrate(&User{}, &Item{}, &UserItem{}, &Quest{})
	fmt.Println("automigration complete")
}

type User struct {
	gorm.Model
	UserName   string
	CurrentHp int
	Xp         int
}

type Item struct {
	gorm.Model
	Name       string
	Attack     int
	Defense    int
}

type UserItem struct {
	gorm.Model
	User       User
	UserID     uint
	Item     Item
	ItemID     uint
}

type Quest struct {
	gorm.Model
	Loction1       string 
	Loction2       string 
	Loction3       string 
	Status         int
	User           User
	UserID         uint
}

