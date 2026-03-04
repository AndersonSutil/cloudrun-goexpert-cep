package handler

import (
	"cloudrun/internal/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type MockWeatherClient struct {
	LocationFunc func(zipcode string) (string, error)
	WeatherFunc  func(city, apiKey string) (float64, error)
}

func (m *MockWeatherClient) GetLocationByZipCode(zipcode string) (string, error) {
	return m.LocationFunc(zipcode)
}

func (m *MockWeatherClient) GetWeatherByCity(city, apiKey string) (float64, error) {
	return m.WeatherFunc(city, apiKey)
}

func TestWeatherHandler_Success(t *testing.T) {
	os.Setenv("WEATHER_API_KEY", "test_key")
	mockClient := &MockWeatherClient{
		LocationFunc: func(zipcode string) (string, error) {
			return "Sao Paulo", nil
		},
		WeatherFunc: func(city, apiKey string) (float64, error) {
			return 28.5, nil
		},
	}
	h := &WeatherHandler{WeatherClient: mockClient}

	req, err := http.NewRequest("GET", "/weather?zipcode=01001000", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response model.WeatherResponse
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if response.TempC != 28.5 {
		t.Errorf("expected 28.5, got %f", response.TempC)
	}
}

func TestWeatherHandler_InvalidZipCode(t *testing.T) {
	h := &WeatherHandler{}
	req, err := http.NewRequest("GET", "/weather?zipcode=123", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnprocessableEntity)
	}
}
