package data

import (
	"encoding/json"
	"io/ioutil"
)

type Response struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	HasBeach    bool      `json:"hasBeach"`
	HasMountain bool      `json:"hasMountain"`
	TempC       []float64 `json:"tempC"`
}

//go:generate mockgen -destination=../mocks/mock_reader.go -package=mocks data.go DataReader
type DataReader interface {
	ReadData() ([]Response, error)
}

type reader struct {
	path string
}

func NewReader() DataReader {
	return &reader{
		path: "./data/cities.json",
	}
}

func (r *reader) ReadData() ([]Response, error) {
	file, err := ioutil.ReadFile(r.path)
	if err != nil {
		return nil, err
	}
	var data []Response
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
