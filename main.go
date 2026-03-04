package main

import (
	"cloudrun/internal/handler"
	"cloudrun/internal/service"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	weatherHandler := &handler.WeatherHandler{
		WeatherClient: &service.DefaultWeatherClient{
			HttpClient: http.DefaultClient,
		},
	}

	http.Handle("/weather", weatherHandler)

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
