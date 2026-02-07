package dotenv

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func parseInt(key string, defaultValue int) (int, error) {
	var value int
	valueString := os.Getenv(key)
	if valueString == "" {
		value = defaultValue
		return value, nil
	}
	value, err := strconv.Atoi(valueString)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %s: %w", valueString, err)
	}

	return value, nil
}

func parseInt64(key string, defaultValue int64) (int64, error) {
	var value int64
	valueString := os.Getenv(key)
	if valueString == "" {
		value = defaultValue
		return value, nil
	}
	value, err := strconv.ParseInt(valueString, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %s: %w", valueString, err)
	}

	return value, nil
}

func parseDuration(key string, defaultValue int) (time.Duration, error) {
	value, err := parseInt(key, defaultValue)
	if err != nil {
		return 0, err
	}

	duration := time.Duration(value) * time.Second
	return duration, nil
}
