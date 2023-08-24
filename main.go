package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	fmt.Println(os.Getenv("PORT"), " ", os.Getenv("API_KEY"))

}
