package core

import (
	"log"
	"os"
)

func workingDir() string {
	workingDir, err := os.Getwd()
	log.Println(err)

	return workingDir
}
