package internal

import (
	"gopkg.in/yaml.v3"
	"os"
)

func ReadRecordsFromFile(filename string) (Records, error) {
	r := Records{}

	data, err := os.ReadFile(filename)
	if err != nil {
		return r, err
	}

	err = yaml.Unmarshal([]byte(data), &r)
	return r, err
}
