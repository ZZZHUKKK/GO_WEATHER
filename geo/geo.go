package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type GeoStruct struct {
	City string `json:"city"`
}

type CityExistance struct {
	Error bool `json:"error"`
}

var ErrNoCity = errors.New("NOCITY")
var ErrNot200 = errors.New("BAD_STATUS_CODE")

func GetMyGeo(city string) (*GeoStruct, error) {
	if city != "" {
		if isExist := CheckCity(city); !isExist {
			return nil, ErrNoCity
		}
		return &GeoStruct{City: city}, nil
	}

	response, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %v", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, ErrNot200
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

func CheckCity(city string) bool {
	postBody, _ := json.Marshal(map[string]string{
		"city": city,
	})
	resp, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	var populationResponce CityExistance
	json.Unmarshal(body, &populationResponce)
	return !populationResponce.Error

	// bodyResp, err := json.Marshal(map[string]string{
	// 	"city": city,
	// })

	// if err != nil {
	// 	return false, errors.New("JSONBAD")
	// }

	// response, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities", "application/json/", bytes.NewBuffer(bodyResp))

	// if err != nil {
	// 	return false, fmt.Errorf("HTTP request failed: %v", err)
	// }

	// defer response.Body.Close()

	// if response.StatusCode != http.StatusOK {
	// 	errorBody, _ := io.ReadAll(response.Body)
	// 	return false, fmt.Errorf("status %d: %s", response.StatusCode, errorBody)
	// }

	// body, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	return false, fmt.Errorf("failed to read body: %v", err)
	// }

	// var cityExist CityExistance

	// if err := json.Unmarshal(body, &cityExist); err != nil {
	// 	return false, fmt.Errorf("JSON parse error: %v", err)
	// }

	// return !cityExist.Error, nil

}
