package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	ExitCodeSuccess = iota
	ExitCodeError
	ExitCodeConfigError
	ExitCodeBackupError
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "BackupConfig.json", "Set a new config to use")
	var deleteLast int
	flag.IntVar(&deleteLast, "d", 0, "Delete the last X backups")
	flag.Parse()

	files, err := GetFilesFromConfig(configFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(ExitCodeConfigError)
	}

	for _, file := range files {
		if deleteLast > 0 {
			DeleteLastXBackups(file, deleteLast)
		}

		if err := CreateBackup(file); err != nil {
			os.Exit(ExitCodeBackupError)
		}
	}

	os.Exit(ExitCodeSuccess)
}