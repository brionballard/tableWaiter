package utils

import (
	"github.com/bxcodec/faker/v4"
	"tableWaiter/restaurant"
	"time"
)

// GenerateRandomTimeStringBetweenOpenAndClose Generates a random time string
// Restaurant opens at 4 p.m. and stops reservations at 10 p.m.
func GenerateRandomTimeStringBetweenOpenAndClose() string {
	startHour := generateRandomNumForTimeString(restaurant.OpensAsInt, restaurant.ReservationsEndAsInt) // Generate random hour between 6 and 10 (inclusive)
	startMinute := generateRandomNumForTimeString(0, 59)                                                // Generate random minute between 0 and 59

	return formatTimeString(startHour, startMinute)
}

func generateRandomNumForTimeString(start int, end int) int {
	num, _ := faker.RandomInt(start, end, 1)
	return num[0]
}

func formatTimeString(a int, b int) string {
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
