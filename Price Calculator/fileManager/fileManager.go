package fileManager

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type FileManager struct {
	InputFileName  string
	OutputFileName string
}

func (fm FileManager) ReadFile() ([]string, error) {
	file, err := os.Open(fm.InputFileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()

	if err != nil {
		fmt.Println("Println error reading file:", err)
		return nil, err
	}
	return lines, nil
}

func (fm FileManager) WriteJSON(data interface{}) error {
	file, err := os.Create(fm.OutputFileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)
	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		return err
	}
	return nil
}

func NewFileManager(inputFileName, outputFileName string) FileManager {
	return FileManager{
		InputFileName:  inputFileName,
		OutputFileName: outputFileName,
	}
}
