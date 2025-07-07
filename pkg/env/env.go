package env

import (
	"fmt"
	"url_shortening/pkg/projectError"

	"os"
	"strconv"
)

func GetEnvOrDie(key string) (string, error) {
	value := os.Getenv(key)

	if value == "" {
		return "", &projectError.Error{
			Code:    projectError.ECONFLICT,
			Message: fmt.Sprintf("Missing environment variable %s", key),
		}
	}

	return value, nil
}

func GetEnvOrDieAsInt(key string) (int, error) {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return -1, &projectError.Error{
			Code:    projectError.ECONFLICT,
			Message: fmt.Sprintf("Environment variable %s not set\n", key),
		}
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return -1, &projectError.Error{
			Code:    projectError.ECONFLICT,
			Message: fmt.Sprintf("Error converting %s to int: %v\n", key, err),
		}
	}

	return value, nil
}
