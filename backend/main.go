package main

import (
	"github.com/MasLazu/CheatChatV2/database"
	"github.com/MasLazu/CheatChatV2/middleware"
	"github.com/MasLazu/CheatChatV2/router"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	database.DBInit()
	defer database.CloseDBConn()

	routerMux := mux.NewRouter()

	apiRoute := routerMux.PathPrefix("/api").Subrouter()

	router.MainRouter(apiRoute)

	log.Println("server runing on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", middleware.EnableCORS(routerMux)))
}
