package rules

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"time"
)

// поменять название пакета
func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	nextDate := ""

	if repeat == "" {
		return "", errors.New("empty repeat rule")
	}
	switch repeat[0] {

	case 'd':

		repeat = repeat[1:]
		daysCount, err := strconv.Atoi(strings.TrimSpace(repeat))

		if err != nil {
			log.Fatal(err)
		}

		if daysCount > 400 {
			log.Fatal("wrong days count is limited by 400")
		}

		dstartTime, err := time.Parse("20060102", dstart)

		if err != nil {
			log.Fatal(err)
		}
		resultTime := time.Time{}
		for {
			resultTime = dstartTime.AddDate(0, 0, daysCount)
			if resultTime.After(now) {
				break
			}
		}

		if !resultTime.After(now) {
			log.Fatal("incorrect result time")
		}

		nextDate = resultTime.Format("20060102")
		break
	case 'y':
		repeat = repeat[1:]
		strings.TrimSpace(repeat)
		if repeat != "" {
			log.Fatal("wrong arguments every year repeats")
		}

		dstartTime, err := time.Parse("20060102", dstart)

		if err != nil {
			log.Fatal(err)
		}
		resultTime := time.Time{}
		for {

			resultTime = dstartTime.AddDate(1, 0, 0)
			if resultTime.After(now) {
				break
			}
		}

		if !resultTime.After(now) {
			log.Fatal("incorrect result time")
		}

		nextDate = resultTime.Format("20060102")

		break
	case 'w':

		break
	case 'm':

		break

	default:
		log.Println("unavailable rule symbol")
		break
	}

	return nextDate, nil
}
