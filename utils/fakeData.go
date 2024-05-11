package utils

import (
	"fmt"
	"github.com/bxcodec/faker/v4"
	"tableWaiter/restaurant"
	"time"
)

// GenerateRandomTimeString Generates a random time string
func GenerateRandomTimeString(minHour int, maxHour int) (string, error) {
	if maxHour > 24 {
		return "", fmt.Errorf("max hour cannot be greater than 24")
	}

	if minHour < 1 {
		return "", fmt.Errorf("min hour cannot be less than 1")
	}

	startHour := GenerateRandomNumForTimeString(minHour, maxHour) // Generate random hour between 6 and 10 (inclusive)
	startMinute := GenerateRandomNumForTimeString(0, 59)          // Generate random minute between 0 and 59

	return FormatTimeString(startHour, startMinute), nil
}

func GenerateRandomNumForTimeString(start int, end int) int {
	num, _ := faker.RandomInt(start, end, 1)
	return num[0]
}

// formatTimeString formats the hour and minute into a string representation of time.
func FormatTimeString(a int, b int) string {
	return time.Date(0, 1, 1, a, b, 0, 0, time.UTC).Format(restaurant.TimeLayout)
}

// GenerateRandomBool generates a random bool
func GenerateRandomBool() bool {
	num, err := faker.RandomInt(restaurant.OpensAsInt, restaurant.ReservationsEndAsInt)
	if err != nil {
		return false
	}
	return num[0] > 7
}

// GetRandomStringFromSlice returns a random string from the given slice.
func GetRandomStringFromSlice(str []string) string {
	selected, err := faker.RandomInt(0, len(str)-1, 1)
	if err != nil {
		fmt.Println(err)
	}

	index := selected[0]
	if selected[0] > len(str) {
		index = len(str)
	}

	return str[index]
}

// GetRandomSBetweenMax generates a random integer between 2 and the given maximum value.
func GetRandomSBetweenMax(max int) int {
	num, err := faker.RandomInt(2, max, 1)
	if err != nil {
		fmt.Println(err)
	}

	return num[0]
}
