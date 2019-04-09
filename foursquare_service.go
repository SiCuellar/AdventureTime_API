package main

import (
	"encoding/json"
	db "github.com/SiCuellar/AdventureTime_API/migrations"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func getQuestLocations(lat string, long string) []FsItem {
	resp, err := http.Get("https://api.foursquare.com/v2/venues/explore?client_id=" + os.Getenv("FOUR_ID") + "&client_secret=" + os.Getenv("FOUR_SECRET") + "&v=20190404&ll=" + lat + "," + long + ",&radius=750&limit=4")

	if err != nil {
		log.Fatal(err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		var result ExploreResult
		json.Unmarshal([]byte(data), &result)
		return result.Response.Groups[0].Items
	}

	return []FsItem{}
}

func buildQuest(lat string, long string, userID uint64) db.Quest {
	locations := getQuestLocations(lat, long)

	quest := db.Quest{
		Location1: locations[0].Venue.Id + "|" + strings.Join(locations[0].Venue.Location.FormattedAddress, ", "),
		Location2: locations[1].Venue.Id + "|" + strings.Join(locations[1].Venue.Location.FormattedAddress, ", "),
		Location3: locations[2].Venue.Id + "|" + strings.Join(locations[2].Venue.Location.FormattedAddress, ", "),
		UserID:    uint(userID)}

	db.NewQuest(quest)

	return quest
}

func snapToLocation(lat string, long string) []string {
	resp, err := http.Get("https://api.foursquare.com/v2/venues/search?client_id=" + os.Getenv("FOUR_ID") + "&client_secret=" + os.Getenv("FOUR_SECRET") + "&v=20190409&ll=" + lat + "," + long + ",&radius=500&limit=3")

	if err != nil {
		log.Fatal(err)
		return nil
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		var result CheckinResult
		json.Unmarshal([]byte(data), &result)

		var ids []string

		for _, venue := range result.Response.Venues {
			ids = append(ids, venue.Id)
		}

		return ids
	}
}

type CheckinResult struct {
	Response struct {
		Venues []Venue `json:"venues"`
	} `json:"response"`
}

type ExploreResult struct {
	Response struct {
		Groups []struct {
			Items []FsItem
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
type FsItem struct {
	Venue Venue
}
