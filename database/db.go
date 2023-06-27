package database

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func checkFile(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		_, err := os.Create(path)
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadDB[T any](path string) ([]T, error) {
	err := checkFile(path)
    if err != nil {
        return nil, err
    }

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if len(content) == 0 {
		return nil, nil
	}

	var data []T

	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func SaveDB[T any](path string, data T) error {
	content, err := json.Marshal(data)
	if err != nil {
		return nil
	}

	err = ioutil.WriteFile(path, content, 0644)
	if err != nil {
		return nil
	}
	return nil
}
