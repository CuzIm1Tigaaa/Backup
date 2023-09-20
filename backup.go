package main

import (
	"fmt"
	Copy "github.com/otiai10/copy"
	"os"
	"path/filepath"
	"sort"
	"time"
)

func CreateBackup(fileName string) error {
	if _, err := os.Stat(fileName); err != nil {
		return err
	}

	fmt.Print("Creating Backup for folder ", fileName, "\n")
	t := time.Now()
	suffix := fmt.Sprintf("%02d.%02d.%v %02d_%02d_%02d",
		t.Day(), int(t.Month()), t.Year(), t.Hour(), t.Minute(), t.Second())

	return Copy.Copy(fileName,
		fmt.Sprintf(".%s%s%s%s %s", string(filepath.Separator), "backups", string(filepath.Separator), fileName, suffix))
}

func DeleteLastXBackups(fileName string, amount int) {
	files, err := filepath.Glob(fmt.Sprintf("backups%s%s *", string(filepath.Separator), fileName))
	if err != nil {
		fmt.Println(err)
		return
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i] < files[j]
	})

	length := amount
	if len(files) < amount {
		length = len(files)
	}

	for i := 0; i < length; i++ {
		err = os.Remove(files[i])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Print("Deleting old Backup ", files[i], "\n")
		}
	}
}

func GetFilesFromConfig(configFile string) ([]string, error) {
	config, err := ParseFilesFromConfig(configFile)
	if err != nil {
		return nil, err
	}
	return config.Files, nil
}