package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type MiRespuesta struct {
    Mensaje string `json:"mensaje"`
}

func main() {
    // Define una función de controlador para el punto final GET
    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
        // Crear una estructura de respuesta JSON
        respuesta := MiRespuesta{Mensaje: "¡Hola, este es un punto final GET en Go!"}

        // Codifica la estructura de respuesta en formato JSON
        jsonRespuesta, err := json.Marshal(respuesta)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Establece el encabezado Content-Type para JSON
        w.Header().Set("Content-Type", "application/json")

        // Escribe la respuesta JSON en el cuerpo de la respuesta HTTP
        w.Write(jsonRespuesta)
    })

    // Inicia el servidor en el puerto 8080
    if err := http.ListenAndServe(":8086", nil); err != nil {
        fmt.Println("Error:", err)
    }
}
