package main

import (
	"log"
	"net/http"

	"github.com/MasLazu/CheatChatV2/app"
	"github.com/MasLazu/CheatChatV2/database"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	database.DBInit()
	defer database.CloseDBConn()

	handler := app.BootstrapApp()

	log.Println("server runing on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", handler))
}
