package nasa

import (
	"encoding/json"
	"io/ioutil"
	"nasawebgo/model"
	"net/http"
)

// GetNasaAPODToday NASA
func GetNasaAPODToday() (model.Nasa, error) {
	var nasaStruct model.Nasa
	resp, err := http.Get("https://api.nasa.gov/planetary/apod?api_key=xpp6bj4qgYwO5shXxXlRxkG2B3MqAyIjzRIrX9Y3")
	if err != nil {
		return nasaStruct, err
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to NASA struct
	err = json.Unmarshal(bodyBytes, &nasaStruct)

	if err != nil {
		return nasaStruct, err
	}

	return nasaStruct, nil
}

// GetNasaAPODS NASA
func GetNasaAPODS() ([]model.Nasa, error) {
	resp, err := http.Get("https://api.nasa.gov/planetary/apod?api_key=xpp6bj4qgYwO5shXxXlRxkG2B3MqAyIjzRIrX9Y3&count=8")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to NASA Array
	var nasaArray = []model.Nasa{}
	err = json.Unmarshal(bodyBytes, &nasaArray)

	if len(nasaArray) == 0 {
		return nil, err
	}

	return nasaArray, nil

}
