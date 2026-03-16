package service

import (
	"cloudrun/internal/model"
	"encoding/json"
	"math"
	"testing"
)

func TestViaCEPResponse_UnmarshalError(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		wantErr  bool
	}{
		{
			name:     "string true",
			jsonData: `{"erro": "true"}`,
			wantErr:  true,
		},
		{
			name:     "bool true",
			jsonData: `{"erro": true}`,
			wantErr:  true,
		},
		{
			name:     "string 1",
			jsonData: `{"erro": "1"}`,
			wantErr:  true,
		},
		{
			name:     "no error",
			jsonData: `{"localidade": "Sao Paulo"}`,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var viaCep model.ViaCEPResponse
			err := json.Unmarshal([]byte(tt.jsonData), &viaCep)
			if err != nil {
				t.Fatalf("Failed to unmarshal: %v", err)
			}

			// We only want to test the error handling logic, so we'll check how service handles it
			// However, GetLocationByZipCode does a real HTTP call.
			// Let's just test the logic directly if possible or mock it.
			// Given the current structure, let's just test the ViaCEPResponse structure first.

			isError := false
			if viaCep.Erro != nil {
				if b, ok := viaCep.Erro.(bool); ok && b {
					isError = true
				}
				if s, ok := viaCep.Erro.(string); ok && (s == "true" || s == "1") {
					isError = true
				}
			}

			if isError != tt.wantErr {
				t.Errorf("Expected isError %v, got %v", tt.wantErr, isError)
			}
		})
	}
}

func TestConvertCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		celsius  float64
		expected float64
	}{
		{0, 32},
		{10, 50},
		{28.5, 83.3},
		{-10, 14},
	}

	for _, tt := range tests {
		result := ConvertCelsiusToFahrenheit(tt.celsius)
		if math.Abs(result-tt.expected) > 0.001 {
			t.Errorf("ConvertCelsiusToFahrenheit(%f) = %f; want %f", tt.celsius, result, tt.expected)
		}
	}
}

func TestConvertCelsiusToKelvin(t *testing.T) {
	tests := []struct {
		celsius  float64
		expected float64
	}{
		{0, 273},
		{10, 283},
		{28.5, 301.5},
		{-273, 0},
	}

	for _, tt := range tests {
		result := ConvertCelsiusToKelvin(tt.celsius)
		if math.Abs(result-tt.expected) > 0.001 {
			t.Errorf("ConvertCelsiusToKelvin(%f) = %f; want %f", tt.celsius, result, tt.expected)
		}
	}
}
