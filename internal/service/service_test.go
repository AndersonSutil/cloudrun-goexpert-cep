package service

import (
	"math"
	"testing"
)

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
