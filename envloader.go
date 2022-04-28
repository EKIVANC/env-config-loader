package envloader

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func LoadEnvVariables(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		currBytes := scanner.Bytes()
		currText := string(currBytes)
		temp := strings.Split(currText, "=")
		if len(temp) != 2 {
			fmt.Fprintf(os.Stderr, "misconfigured env variable, pls check app.config file")
			os.Exit(1)
		}
		key := temp[0]
		value := temp[1]
		_, exists := os.LookupEnv(key)
		// only set env variable if it does not already exist, so container level env config always overwrites app.config file settings
		if !exists {
			err := os.Setenv(key, value)
			if err != nil {
				fmt.Fprintf(os.Stderr, "environment variable could not set: %s , error: %v \n", key, err)
				os.Exit(1)
			}
		}
	}
}
