package utils

import (
	"os"
	"strconv"
)

func GetStringEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetIntEnv(key string, fallback int) int {
	if strValue, ok := os.LookupEnv(key); ok {
		value, err := strconv.Atoi(strValue)
		if err != nil {
			return fallback
		}
		return value
	}
	return fallback
}

func GetBoolEnv(key string, fallback bool) bool {
	if strValue, ok := os.LookupEnv(key); ok {
		value, err := strconv.ParseBool(strValue)
		if err != nil {
			return fallback
		}
		return value
	}
	return fallback
}
