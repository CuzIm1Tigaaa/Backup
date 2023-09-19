package main

import (
	"fmt"
	Copy "github.com/otiai10/copy"
	"os"
	"time"
)

func CreateBackup(fileName string) error {
	if _, err := os.Stat(fileName); err != nil {
		return err
	}

	t := time.Now()
	suffix := fmt.Sprintf("%02d.%02d.%v %02d_%02d_%02d",
		t.Day(), int(t.Month()), t.Year(), t.Hour(), t.Minute(), t.Second())

	return Copy.Copy(fileName,
		fmt.Sprintf("%s/%s %s", "backups", fileName, suffix))
}

func DeleteLastXBackups(amount int) {

}

func GetFilesFromConfig(configFile string) ([]string, error) {
	config, err := ParseFilesFromConfig(configFile)
	if err != nil {
		return nil, err
	}
	return config.Files, nil
}