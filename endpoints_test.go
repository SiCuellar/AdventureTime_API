package main

import (
	"encoding/json"
	"github.com/SiCuellar/AdventureTime_API/migrations"
	"github.com/jinzhu/gorm"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	db_url := "\n host=localhost\n port=5432\n user=postgres\n dbname=adventuretime_test\n sslmode=disable\n password="

	// fmt.Printf("Using database config string: %s\n", db_url)

	db.Connection, _ = gorm.Open("postgres", db_url)
	ResetDatabase()
}

func ResetDatabase() {
	db.Connection.DropTable(&db.User{}, &db.Item{}, &db.UserItem{}, &db.Quest{})
	db.Connection.AutoMigrate(&db.User{}, &db.Item{}, &db.UserItem{}, &db.Quest{})

	user := db.User{UserName: "test", Xp: 200, CurrentHp: 100}

	db.Connection.NewRecord(user)
	db.Connection.Create(&user)
}

func TestRootHandler(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	Router().ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code, "Status 200: OK expected.")
	assert.Equal(t, "Hello World!\n", response.Body.String(), "Expect body of \"Hello World!\"")
}

func TestLoginHandler(t *testing.T) {
	defer ResetDatabase()
	request, _ := http.NewRequest("POST", "/api/v1/login?username=test", nil)
	response := httptest.NewRecorder()

	Router().ServeHTTP(response, request)

	type APIResponse struct {
		Value db.User `json:"Value"`
	}

	var res APIResponse

	json.Unmarshal(response.Body.Bytes(), &res)

	assert.Equal(t, 200, response.Code, "Status 200: OK expected.")
	assert.Equal(t, "test", res.Value.UserName, "Expected json response to return User Object")
}

