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
	dbURL := os.Getenv("DATABASE_URL")

	if dbURL == "" {
		dbURL = "\n host=localhost\n port=5432\n user=postgres\n dbname=adventuretime\n sslmode=disable\n password="
	}

	// fmt.Printf("Using database config string: %s\n", dbURL)
	Connection, err = gorm.Open("postgres", dbURL)

	if err != nil {
		fmt.Println(err)
	}
	// defer fmt.Println("Connected to db adventuretime")
}

func Migrate() {
	Connect()
	Connection.AutoMigrate(&User{}, &Item{}, &UserItem{}, &Quest{})
	defer fmt.Println("Automigration Complete.")
}

func Close() {
	Connection.Close()
	// defer fmt.Println("Connection to db adventuretime closed.")
}

func NewQuest(newQuest Quest) {
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
	User   User `json:"-"`
	UserID uint `json:"user_id"`
	Item   Item `json:"-"`
	ItemID uint `json:"item_id"`
}

type Quest struct {
	gorm.Model
	Location1 string `json:"location_1"`
	Location2 string `json:"location_2"`
	Location3 string `json:"location_3"`
	Status    int    `json:"status"`
	User      User   `json:"-"`
	UserID    uint   `json:"user_id"`
}
