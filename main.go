package main

import (
	"github.com/SiCuellar/AdventureTime_API/migrations"
	// "github.com/gorilla/mux"
)

func main() {
	migrations.Migrate()
	
}
