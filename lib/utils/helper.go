package utils

import "github.com/google/uuid"

func GetUUID() string {
	uuid, _ := uuid.NewUUID()

	return uuid.String()
}

func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
