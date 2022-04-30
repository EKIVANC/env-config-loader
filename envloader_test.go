package envloader

import (
	"os"
	"testing"
)

// test if env param set correctly
func TestLoadEnvVariables(t *testing.T) {
	testEnvParam := "TEST_PARAM"
	LoadEnvVariables("./configs/app_test.config")
	if _, ok := os.LookupEnv(testEnvParam); !ok {
		t.Errorf("The parameter %s not found in env variables", testEnvParam)
	}
}

// Test if env parameters set after a comment line which starts with #
func TestCommentLines(t *testing.T) {
	testEnvParam := "TEST_PARAM_2"
	LoadEnvVariables("./configs/app_test.config")
	if _, ok := os.LookupEnv(testEnvParam); !ok {
		t.Errorf("The parameter %s not found in env variables", testEnvParam)
	}
}

// test if blank lines breaks env parameters setting or not
func TestBlankLines(t *testing.T) {
	testEnvParam := "TEST_PARAM_3"
	LoadEnvVariables("./configs/app_test.config")
	if _, ok := os.LookupEnv(testEnvParam); !ok {
		t.Errorf("The parameter %s not found in env variables", testEnvParam)
	}
}

func TestTrimmingWhiteSpaces(t *testing.T) {
	LoadEnvVariables("./configs/app_test.config")
	val := os.Getenv("TEST_PARAM_4")
	if val != "TEST_VALUE_4" {
		t.Errorf("Trimming feature failed for env variable or its value")
	}
}
