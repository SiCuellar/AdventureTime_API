package main

import (
	"fmt"
	db "github.com/SiCuellar/AdventureTime_API/migrations"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func main() {
	db.Migrate()
	db.Connect()

	router := Router()

	mux := cors.Default().Handler(router)

	fmt.Println("Listening on port: " + os.Getenv("PORT"))

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), mux))
}
