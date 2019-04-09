package main

import (
	"encoding/json"
	"fmt"
	db "github.com/SiCuellar/AdventureTime_API/migrations"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
	"strconv"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", RootHandler).Methods("GET")
	router.HandleFunc("/api/v1/login", LoginHandler).Methods("POST")
	router.HandleFunc("/api/v1/quest", QuestHandler).Methods("POST")

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

func CheckParams(params url.Values) bool {
	var latMissing bool
	var longMissing bool

	latMissing = params["lat"] == nil || params["lat"][0] == ""
	longMissing = params["long"] == nil || params["long"][0] == ""

	return latMissing || longMissing
}

type ErrorJSON struct {
	Error string
}
