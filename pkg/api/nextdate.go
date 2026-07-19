package api

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const dataFormat = "20060102"

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

		resultTime, err := time.Parse(dataFormat, dstart)

		if err != nil {
			log.Fatal(err)
			return "", err
		}

		for {
			resultTime = resultTime.AddDate(0, 0, daysCount)
			if resultTime.After(now) {
				break
			}
		}

		if !resultTime.After(now) {
			log.Fatal("incorrect result time")
		}

		nextDate = resultTime.Format(dataFormat)
		break
	case 'y':
		repeat = repeat[1:]
		strings.TrimSpace(repeat)
		if repeat != "" {
			log.Fatal("wrong arguments every year repeats")
		}

		resultTime, err := time.Parse(dataFormat, dstart)

		if err != nil {
			log.Fatal(err)
			return "", err
		}

		for {

			resultTime = resultTime.AddDate(1, 0, 0)
			if resultTime.After(now) {
				break
			}
		}

		if !resultTime.After(now) {
			log.Fatal("incorrect result time")
		}

		nextDate = resultTime.Format(dataFormat)

		break
	case 'w':

		break
	case 'm':

		break

	default:
		return "", errors.New("invalid rule format")
	}

	return nextDate, nil
}

func HandleNextDate(res http.ResponseWriter, req *http.Request) {

	nowValue := req.FormValue("now")
	dateValue := req.FormValue("date")
	repeatValue := req.FormValue("repeat")

	now, err := time.Parse(dataFormat, nowValue)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	nextDate, err := NextDate(now, dateValue, repeatValue)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	res.Write([]byte(nextDate))

	//"api/nextdate?now=20240126&date=20240126&repeat=y"
}
