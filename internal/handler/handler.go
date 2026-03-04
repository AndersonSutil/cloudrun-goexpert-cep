package handler

import (
	"cloudrun/internal/model"
	"cloudrun/internal/service"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"regexp"
)

type WeatherHandler struct {
	WeatherClient service.WeatherClient
}

func (h *WeatherHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	zipcode := r.URL.Query().Get("zipcode")
	if zipcode == "" {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	// Validate zipcode format (8 digits)
	matched, _ := regexp.MatchString(`^\d{8}$`, zipcode)
	if !matched {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		http.Error(w, "internal server error: missing api key", http.StatusInternalServerError)
		return
	}

	city, err := h.WeatherClient.GetLocationByZipCode(zipcode)
	if err != nil {
		if errors.Is(err, service.ErrInvalidZipCode) {
			http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
			return
		}
		if errors.Is(err, service.ErrZipCodeNotFound) {
			http.Error(w, "can not find zipcode", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tempC, err := h.WeatherClient.GetWeatherByCity(city, apiKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := model.WeatherResponse{
		TempC: tempC,
		TempF: service.ConvertCelsiusToFahrenheit(tempC),
		TempK: service.ConvertCelsiusToKelvin(tempC),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
