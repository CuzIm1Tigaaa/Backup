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

func Test_Folder_GetFiles(t *testing.T) {
	// r := require.New(t)

}