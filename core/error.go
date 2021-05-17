package core

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const printError = false

// Fancy log function
func Error(info string, err error) {
	if err == nil {
		return
	}

	_, fn, line, _ := runtime.Caller(1)

	path := strings.Split(fn, string(filepath.Separator))
	fn = path[len(path)-1]

	if printError {
		fmt.Printf("%v\n\n\n[Error] %s in %s, Line: %d\n\n", err, info, fn, line)
	} else {
		fmt.Printf("[Error] %s in %s, Line: %d\n\n", info, fn, line)
	}

	logSession(err.Error())

	os.Exit(0)
}

func logSession(errStr string) {
	f, err := os.OpenFile(workingDir()+"session.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	log.Println(err)
	defer f.Close()

	if _, err := f.WriteString(errStr); err != nil {
		log.Println(err)
	}
}
