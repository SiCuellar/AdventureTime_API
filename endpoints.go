package main

import (
	"encoding/json"
	"fmt"
	db "github.com/SiCuellar/AdventureTime_API/migrations"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", RootHandler).Methods("GET")
	router.HandleFunc("/api/v1/login", LoginHandler).Methods("POST")
	router.HandleFunc("/api/v1/quest", QuestHandler).Methods("POST")
	router.HandleFunc("/api/v1/checkin", CheckinHandler).Methods("POST")
	router.HandleFunc("/api/v1/encounter", EncounterHandler).Methods("POST")

	return router
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "Hello World!")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	user := db.Connection.Preload("Items").Where("user_name = ?", params["username"]).First(&db.User{})
	_ = json.NewEncoder(w).Encode(user)
}

func QuestHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	if CheckParams(params) {
		w.WriteHeader(406)
		_ = json.NewEncoder(w).Encode(ErrorJSON{"You must provide a lat and long"})
		return
	}

	userID, _ := strconv.ParseUint(params["user_id"][0], 10, 32)
	lat := params["lat"][0]
	long := params["long"][0]

	var oldQuest db.Quest

	query := db.Connection.Where("status = ?", 0).Where("user_id = ?", userID).First(&oldQuest)

	var quest db.Quest

	if query.RecordNotFound() {
		fmt.Println("Previous Quest not found. Generating new quest.")
		quest = buildQuest(lat, long, userID)
	} else {
		fmt.Println("Previous Quest Found! ")
		quest = oldQuest
	}

	_ = json.NewEncoder(w).Encode(quest)
}

func EncounterHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	userID, _ := strconv.ParseUint(params["user_id"][0], 10, 32)
	var oldQuest db.Quest
	_ = db.Connection.Where("status = ?", 0).Where("user_id = ?", userID).First(&oldQuest)

	
	var user db.User
	db.Connection.First(&user, userID)
	
	if params["success"][0] == "true" {
		oldQuest.CurrentLocation++
		user.Xp+= 100

		var questComplete bool = false

		if oldQuest.CurrentLocation > 3 {
			oldQuest.Status = 2
			questComplete = true
		}

		_ = json.NewEncoder(w).Encode(struct{
			Success string `json:"success"`
			QuestComplete bool `json:"quest_complete"`
		}{"Successful Encounter", questComplete})
	} else {
		_ = json.NewEncoder(w).Encode(ErrorJSON{"Encounter Failed"})
		user.CurrentHp = 10
		user.Xp = 0
		oldQuest.Status = 1 
	}
	db.Connection.Save(&user)
	db.Connection.Save(&oldQuest)
}

func CheckinHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	if CheckParams(params) {
		w.WriteHeader(406)
		_ = json.NewEncoder(w).Encode(ErrorJSON{"You must provide a lat and long"})
		return
	}

	userID, _ := strconv.ParseUint(params["user_id"][0], 10, 32)
	lat := params["lat"][0]
	long := params["long"][0]

	var currentQuest db.Quest

	db.Connection.Where("status = ?", 0).Where("user_id = ?", userID).First(&currentQuest)

	var currentLocationID string

	switch locationIndex := currentQuest.CurrentLocation; locationIndex {
	case 1:
		currentLocationID = strings.Split(currentQuest.Location1, "|")[0]
	case 2:
		currentLocationID = strings.Split(currentQuest.Location2, "|")[0]
	case 3:
		currentLocationID = strings.Split(currentQuest.Location3, "|")[0]
	}

	ids := snapToLocation(lat, long)
	flag := false

	for _, id := range ids {
		if id == currentLocationID {
			flag = true
		}
	}

	if flag {
		_ = json.NewEncoder(w).Encode(SuccessJSON{"Lat/Long matches current goal location."})
	} else {
		_ = json.NewEncoder(w).Encode(ErrorJSON{"Lat/Long does not match current goal location."})
	}
}

func CheckParams(params url.Values) bool {
	var latMissing bool
	var longMissing bool

	latMissing = params["lat"] == nil || params["lat"][0] == ""
	longMissing = params["long"] == nil || params["long"][0] == ""

	return latMissing || longMissing
}

type ErrorJSON struct {
	Error string `json:"error"`
}

type SuccessJSON struct {
	Success string `json:"success"`
}
