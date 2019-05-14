package helper

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Confirm generate confirmation message
func Confirm(message string) (ok bool) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message)
	answer, _ := reader.ReadString('\n')
	answer = strings.Trim(answer, "\n")
	ok, _ = regexp.MatchString("(?i)^y(es)?$", answer)
	if ok {
		fmt.Println("Aborting...")
		os.Exit(0)
	}
	return
}

// CreateDir create directory
func CreateDir(path string, mode os.FileMode, recursive bool) (err error) {
	if recursive {
		err = os.MkdirAll(path, mode)
	} else {
		err = os.Mkdir(path, mode)
	}
	return
}

func CreateFile(name, content string) {
	file, err := os.Create(name)
	defer file.Close()
	if err != nil {
		fmt.Printf("Failed to create file in path %s\n", name)
		os.Exit(1)
	}
	if _, err := file.WriteString(content); err != nil {
		fmt.Printf("Error occurs: %v", err)
		os.Exit(1)
	}
}
