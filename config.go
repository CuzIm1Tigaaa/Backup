package main

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"strings"
)

type Config struct {
	Files []string `json:"files"`
}

func ParseFilesFromConfig(fileName string) (*Config, error) {
	if !strings.HasSuffix(fileName, ".json") {
		return nil, errors.New("the given file is not a json file")
	}
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var config *Config
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}