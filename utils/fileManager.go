package utils

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func FileManager(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error opening file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error reading file")
	}

	err = file.Close()
	if err != nil {
		return nil, errors.New("error closing file")
	}
	return lines, nil
}

func WriteJSON(data any, path string) error {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return errors.New("error creating file")
	}
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("error writing to file")
	}
	err = file.Close()
	if err != nil {
		return errors.New("error closing file")
	}
	return nil
}
