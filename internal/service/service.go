package service

import (
	"cloudrun/internal/model"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

var (
	ErrZipCodeNotFound = errors.New("can not find zipcode")
	ErrInvalidZipCode  = errors.New("invalid zipcode")
)

type WeatherClient interface {
	GetLocationByZipCode(zipcode string) (string, error)
	GetWeatherByCity(city, apiKey string) (float64, error)
}

type DefaultWeatherClient struct {
	HttpClient *http.Client
}

func (c *DefaultWeatherClient) GetLocationByZipCode(zipcode string) (string, error) {
	if len(zipcode) != 8 {
		return "", ErrInvalidZipCode
	}

	httpClient := c.HttpClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	resp, err := httpClient.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", zipcode))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", ErrZipCodeNotFound
	}

	var viaCep model.ViaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&viaCep); err != nil {
		return "", err
	}

	if viaCep.Erro != nil {
		if b, ok := viaCep.Erro.(bool); ok && b {
			return "", ErrZipCodeNotFound
		}
		if s, ok := viaCep.Erro.(string); ok && (s == "true" || s == "1") {
			return "", ErrZipCodeNotFound
		}
	}

	return viaCep.Localidade, nil
}

func (c *DefaultWeatherClient) GetWeatherByCity(city, apiKey string) (float64, error) {
	cityEncoded := url.QueryEscape(city)
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", apiKey, cityEncoded)

	httpClient := c.HttpClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("weather api error: %d", resp.StatusCode)
	}

	var weather model.WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return 0, err
	}

	return weather.Current.TempC, nil
}

func ConvertCelsiusToFahrenheit(c float64) float64 {
	return c*1.8 + 32
}

func ConvertCelsiusToKelvin(c float64) float64 {
	return c + 273
}
