package main

import (
	"log"
	"os"
	"test-wgs/api"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "local" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file not FOUND")
		}
	}
	r := gin.Default()

	api.Init(r)
}
