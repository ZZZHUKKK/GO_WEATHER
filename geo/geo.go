package geo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GeoStruct struct {
	City string `json:"city"`
}

func GetMyGeo(city string) (*GeoStruct, error) {
	if city != "" {
		return &GeoStruct{City: city}, nil
	}

	response, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %v", err)
	}
	defer response.Body.Close() // Закрываем тело в любом случае

	// Проверяем статус ДО чтения тела
	if response.StatusCode != http.StatusOK {
		errorBody, _ := io.ReadAll(response.Body) // Читаем тело ошибки
		return nil, fmt.Errorf("status %d: %s", response.StatusCode, errorBody)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %v", err)
	}

	var geo GeoStruct

	if err := json.Unmarshal(body, &geo); err != nil {
		return nil, fmt.Errorf("JSON parse error: %v", err)
	}

	return &geo, nil
}
