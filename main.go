package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type MiRespuesta struct {
	Mensaje string `json:"mensaje"`
}

func main() {
	Config()

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {

		respuesta := MiRespuesta{Mensaje: "¡Hola, este es un punto final GET en Go!"}

		jsonRespuesta, err := json.Marshal(respuesta)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("key!:", os.Getenv("API_KEY"))

		w.Header().Set("Content-Type", "application/json")

		w.Write(jsonRespuesta)
	})
	PORT := os.Getenv("PORT")
	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		fmt.Println("Error:", err)
	}
}

func Config() {
	err := godotenv.Load("./.env")
	if err != nil {
		_ = godotenv.Load("/go/bin/.env")
	}
}
