package main

import (
	"companies/routers"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	godotenv.Load()

	log.Println("Starting server after loading env vars")
	r := routers.SetupRouter()

	r.Run() // listens on 0.0.0.0:8080 by default
}
