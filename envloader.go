package envloader

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const commentPrefix = "#"

func LoadEnvVariables(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		currBytes := scanner.Bytes()
		currText := string(currBytes)
		currText = strings.TrimSpace(currText)
		if len(currText) == 0 {
			// skip empty lines
			continue
		}
		if strings.HasPrefix(currText, commentPrefix) {
			// skip comment lines
			continue
		}
		temp := strings.Split(currText, "=")
		if len(temp) != 2 {
			fmt.Fprintf(os.Stderr, "misconfigured env variable, pls check %s file , line: %s", path, currText)
			os.Exit(1)
		}
		key := strings.TrimSpace(temp[0])
		value := strings.TrimSpace(temp[1])
		if len(key) == 0 || len(value) == 0 {
			fmt.Fprintf(os.Stderr, "misconfigured env variable, pls check %s file for variable %s %s ", path, key, value)
			os.Exit(1)
		}
		_, exists := os.LookupEnv(key)
		// only set env variable if it does not already exist, so container level env config always overwrites app.config file settings
		if !exists {
			err := os.Setenv(key, value)
			if err != nil {
				fmt.Fprintf(os.Stderr, "environment variable could not set: %s , error: %v file %s \n", key, err, path)
				os.Exit(1)
			}
		}
	}
}
