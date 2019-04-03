package migrations

import (
		"fmt"
		"os"
		"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Migrate() {
	db_url := os.Getenv("DATABASE_URL")

	if db_url == "" {
		db_url = "host=localhost port=5432 user=root dbname=adventuretime sslmode=disable password="
	}

	fmt.Printf("Using database config string: %s\n", db_url)

	db, err := gorm.Open("postgres", db_url)
	defer db.Close()

	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&User{}, &Item{}, &UserItem{}, &Quest{})
	defer fmt.Println("Automigration Complete.")
}

type User struct {
	gorm.Model
	UserName  string
	CurrentHp int
	Xp        int
}

type Item struct {
	gorm.Model
	Name    string
	Attack  int
	Defense int
}

type UserItem struct {
	gorm.Model
	User   User
	UserID uint
	Item   Item
	ItemID uint
}

type Quest struct {
	gorm.Model
	Loction1 string
	Loction2 string
	Loction3 string
	Status   int
	User     User
	UserID   uint
}
