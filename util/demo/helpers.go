package main

import (
	"math/rand"
	"time"
)

func randomString(options []string) string {
	return options[rand.Intn(len(options))]
}

func randomInt(options []int) int {
	return options[rand.Intn(len(options))]
}

// func randomInts(count int, options []int) []int {
// 	ints := []int{}
// 	for i := 0; i < count; i++ {
// 		ints = append(ints, options[rand.Intn(len(options))])
// 	}
// 	return ints
// }

func randomDOBAndAge() (string, int) {
	currentYear := time.Now().Year()

	yo := rand.Intn(50) + 20

	year := currentYear - yo
	month := rand.Intn(12) + 1
	day := rand.Intn(25) + 1

	dob := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC).String()

	return dob, yo
}

func chunkBy(items []string, chunkSize int) (chunks [][]string) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}

	return append(chunks, items)
}
