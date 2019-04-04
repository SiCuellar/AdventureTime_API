package db

import (
		"fmt"
		"os"
		"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var connection *gorm.DB

func Connect() {
	var err error
	db_url := os.Getenv("DATABASE_URL")

	if db_url == "" {
		db_url = "\n host=localhost\n port=5432\n user=root\n dbname=adventuretime\n sslmode=disable\n password="
	}
	
	// fmt.Printf("Using database config string: %s\n", db_url)

	connection, err = gorm.Open("postgres", db_url)

	if err != nil {
		fmt.Println(err)
	}
	// defer fmt.Println("Connected to db adventuretime")
}

func Migrate() {
	Connect()
	connection.AutoMigrate(&User{}, &Item{}, &UserItem{}, &Quest{})
	Close()
	defer fmt.Println("Automigration Complete.")
}

func Close() {
	connection.Close()
	// defer fmt.Println("Connection to db adventuretime closed.")
}

func NewQuest(newQuest Quest) {
	Connect()
	connection.NewRecord(newQuest)
	connection.Create(&newQuest)
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
	Location1 string
	Location2 string
	Location3 string
	Status   int 
	User     User
	UserID   uint
}