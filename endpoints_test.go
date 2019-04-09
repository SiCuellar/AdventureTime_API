package main

import (
	"encoding/json"
	"fmt"
	"github.com/SiCuellar/AdventureTime_API/migrations"
	"github.com/jinzhu/gorm"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init()  {
	SetupDatabase()
}

func SetupDatabase() {
	db_url := "\n host=localhost\n port=5432\n user=postgres\n dbname=adventuretime_test\n sslmode=disable\n password="

	// fmt.Printf("Using database config string: %s\n", db_url)

	db.Connection, _ = gorm.Open("postgres", db_url)
	ResetDatabase()
}

var user db.User

func ResetDatabase() {
	db.Connection.DropTable(&db.User{}, &db.Item{}, &db.UserItem{}, &db.Quest{})
	db.Connection.AutoMigrate(&db.User{}, &db.Item{}, &db.UserItem{}, &db.Quest{})

	user = db.User{UserName: "test", Xp: 200, CurrentHp: 100}

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

func TestQuestHandlerWithNewQuest(t *testing.T) {
	defer ResetDatabase()

	url := fmt.Sprintf("/api/v1/quest?lat=39.7508479&long=-104.9964413&user_id=%v", user.ID)

	request, _ := http.NewRequest("POST", url, nil)
	response := httptest.NewRecorder()

	Router().ServeHTTP(response, request)

	var res db.Quest

	json.Unmarshal(response.Body.Bytes(), &res)

	assert.Equal(t, 200, response.Code, "Status 200: OK expected.")
	assert.Equal(t, LastUser().ID, res.UserID, "Expected quest object to have correct user association")
	assert.Equal(t, "4a8b585ff964a520360c20e3|1822 Blake St (btwn 19th St \u0026 18th St), Denver, CO 80202, United States", res.Location1, "Expected quest object to have 3 locations: location_1")
	assert.Equal(t, "4e10f909483bee47ff2e50c0|Denver, CO 80204, United States", res.Location2, "Expected quest object to have 3 locations: location_2")
	assert.Equal(t, "4a58056ff964a52042b71fe3|1539 17th St (17th and Wazee), Denver, CO 80202, United States", res.Location3, "Expected quest object to have 3 locations: location_3")
	assert.Equal(t, 0, res.Status, "Expected quest object to have 3 locations: location_3")
}

func TestQuestHandlerWithOldQuest(t *testing.T) {
	quest := buildQuest("39.7508479", "-104.9964413", uint64(user.ID))

	db.Connection.NewRecord(quest)
	db.Connection.Create(&quest)

	defer ResetDatabase()

	url := fmt.Sprintf("/api/v1/quest?lat=39.7508479&long=-104.9964413&user_id=%v", user.ID)

	request, _ := http.NewRequest("POST", url, nil)
	response := httptest.NewRecorder()

	Router().ServeHTTP(response, request)

	var res db.Quest

	json.Unmarshal(response.Body.Bytes(), &res)

	assert.Equal(t, 200, response.Code, "Status 200: OK expected.")
	assert.Equal(t, LastUser().ID, res.UserID, "Expected quest object to have a location1")
	assert.Equal(t, quest.Location1, res.Location1, "Expected quest object to have a lcoat")
	assert.Equal(t, quest.Location2, res.Location2, "Expected quest object to have a lcoat")
	assert.Equal(t, quest.Location3, res.Location3, "Expected quest object to have a lcoat")
	assert.Equal(t, 0, res.Status, "Expected quest object to have correct user association")
}

func TestQuestHandlerWithNoLatLong(t *testing.T) {
	defer ResetDatabase()

	url := fmt.Sprintf("/api/v1/quest?lat&long&user_id=%v", user.ID)

	request, _ := http.NewRequest("POST", url, nil)
	response := httptest.NewRecorder()

	Router().ServeHTTP(response, request)

	assert.Equal(t, 406, response.Code, "Status 406 Expected.")
	assert.Equal(t, "{\"Error\":\"You must provide a lat and long\"}\n", response.Body.String(), "Expected Error JSON")
}

func LastUser() db.User {
	var user1 db.User
	db.Connection.Last(&user1)
	return user1
}
