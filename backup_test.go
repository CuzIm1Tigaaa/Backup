package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_CreateBackup(t *testing.T) {
	r := require.New(t)

	file := "./testData/Folder"
	err := CreateBackup(file)
	r.NoError(err, "Backup could not be created!")
}

func Test_GetFilesFromConfig(t *testing.T) {
	r := require.New(t)

	testFile := "./testData/BackupConfig.json"
	files, err := GetFilesFromConfig(testFile)

	r.NoError(err)
	r.Equal(3, len(files))
}

func Test_CreateBackupsFromConfig(t *testing.T) {
	r := require.New(t)

	testFile := "./testData/BackupConfig.json"
	files, err := GetFilesFromConfig(testFile)

	r.NoError(err)

	for _, file := range files {
		err := CreateBackup(file)
		r.NoError(err, "Backup could not be created!")
	}
}

func Test_DeleteLastBackups(t *testing.T) {
	r := require.New(t)

	testFile := "./testData/BackupConfig.json"
	files, err := GetFilesFromConfig(testFile)

	r.NoError(err)

	for _, file := range files {
		DeleteLastXBackups(file, 1)
	}
}