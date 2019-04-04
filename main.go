package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

  "github.com/SiCuellar/AdventureTime_API/environment"
  "github.com/SiCuellar/AdventureTime_API/migrations"
)

func main() {
	db.Migrate()
	// environment.SetVariables()
	// buildQuest()

	db.Connect()

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/login", GetLoginHandler).Methods("POST")

	fmt.Println("Listening on port: " + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}

func GetLoginHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	user := db.Connection.Preload("Items").Find(&db.User{}, params["user_id"])
	json.NewEncoder(w).Encode(user)
}

func getQuestLocations() []Item {
	resp, err := http.Get("https://api.foursquare.com/v2/venues/explore?client_id=" + os.Getenv("FOUR_ID") + "&client_secret=" + os.Getenv("FOUR_SECRET") + "&v=20190401&ll=39.7527044,-104.9918035,&radius=100")
	if err != nil {
		log.Fatal(err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		var result Result
		json.Unmarshal([]byte(data), &result)
		return result.Response.Groups[0].Items[0:3]
	}
	return []Item{}
}

func buildQuest() {
	for _, item := range getQuestLocations() {
		locations := item.Venue.Location.FormattedAddress
		quest := db.Quest{Location1: locations[0], Location2: locations[1], Location3: locations[2]}
		db.NewQuest(quest)
	}
	defer db.Close()
}

type Result struct {
	Response struct {
		Groups []struct {
			Items []Item
		}
	}
}

type Location struct {
	Lat              float64
	Lng              float64
	Distance         int
	FormattedAddress []string // check this later
}

type Venue struct {
	Id       string
	Name     string
	Location Location
}

type Item struct {
	Venue Venue
}
