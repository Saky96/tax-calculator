package utils

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type FileManager struct {
	InputSource  string
	OutputSource string
}

func NewFileManager(inputSource, outputSource string) FileManager {
	return FileManager{
		InputSource:  inputSource,
		OutputSource: outputSource,
	}
}

func (fm FileManager) ReadFile() ([]string, error) {
	file, err := os.Open(fm.InputSource)
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

func (fm FileManager) WriteFile(data any) error {
	file, err := os.Create(fm.OutputSource)
	if err != nil {
		fmt.Println(err)
		return errors.New("error creating file")
	}

	time.Sleep(2 * time.Second) // Simulate some delay for demonstration purposes

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
