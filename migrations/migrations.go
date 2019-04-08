package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Connection *gorm.DB

func Connect() {
	var err error
	db_url := os.Getenv("DATABASE_URL")

	if db_url == "" {
		db_url = "\n host=localhost\n port=5432\n user=postgres\n dbname=adventuretime\n sslmode=disable\n password="
	}

	// fmt.Printf("Using database config string: %s\n", db_url)

	Connection, err = gorm.Open("postgres", db_url)

	if err != nil {
		fmt.Println(err)
	}
	// defer fmt.Println("Connected to db adventuretime")
}

func Migrate() {
	Connect()
	Connection.AutoMigrate(&User{}, &Item{}, &UserItem{}, &Quest{})
	Close()
	defer fmt.Println("Automigration Complete.")
}

func Close() {
	Connection.Close()
	// defer fmt.Println("Connection to db adventuretime closed.")
}

func NewQuest(newQuest Quest) {
	Connect()
	Connection.NewRecord(newQuest)
	Connection.Create(&newQuest)
}

type User struct {
	gorm.Model
	UserName  string `json:"username"`
	CurrentHp int    `json:"current_hp" gorm:"default: 10"`
	Xp        int    `json:"current_xp"`
	Items     []Item `gorm:"many2many:user_items" json:"items"`
}

type Item struct {
	gorm.Model
	Name    string `json:"name"`
	Attack  int    `json:"attack"`
	Defense int    `json:"defense"`
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
	Location1 string ``
	Location2 string ``
	Location3 string ``
	Status    int
	User      User `json:"-"`
	UserID    uint
}
