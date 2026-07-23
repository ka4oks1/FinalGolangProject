package api

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const dateFormat = "20060102"

func afterNow(date, now time.Time) bool {

	return date.After(now)
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	nextDate := ""

	if repeat == "" {
		return "", nil
	}

	if dstart == "" {
		dstart = time.Now().Format(dateFormat)
	}

	switch repeat[0] {

	case 'd':

		repeat = repeat[1:]
		daysCount, err := strconv.Atoi(strings.TrimSpace(repeat))

		if err != nil {
			return "", err
		}

		if daysCount > 400 {
			return "", err
		}

		resultTime, err := time.Parse(dateFormat, dstart)

		if err != nil {
			//log.Fatal(err)
			return "", err
		}

		for {
			resultTime = resultTime.AddDate(0, 0, daysCount)
			if afterNow(resultTime, now) {
				break
			}
		}

		if !afterNow(resultTime, now) {
			return "", err
		}

		nextDate = resultTime.Format(dateFormat)
		break
	case 'y':
		repeat = repeat[1:]
		strings.TrimSpace(repeat)
		if repeat != "" {
			log.Fatal("wrong arguments every year repeats")
		}

		resultTime, err := time.Parse(dateFormat, dstart)

		if err != nil {
			return "", err
		}

		for {

			resultTime = resultTime.AddDate(1, 0, 0)
			if afterNow(resultTime, now) {
				break
			}
		}

		if !afterNow(resultTime, now) {
			return "", err
		}

		nextDate = resultTime.Format(dateFormat)

		break
	case 'w':
		return "", errors.New("invalid rule format")
	case 'm':
		return "", errors.New("invalid rule format")
	default:
		return "", errors.New("invalid rule format")
	}

	return nextDate, nil
}

func HandleNextDate(res http.ResponseWriter, req *http.Request) {

	nowValue := req.FormValue("now")
	dateValue := req.FormValue("date")
	repeatValue := req.FormValue("repeat")

	now, err := time.Parse(dateFormat, nowValue)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	nextDate, err := NextDate(now, dateValue, repeatValue)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	res.Write([]byte(nextDate))

}
