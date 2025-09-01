package main

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"
)

var IncorrectInput = errors.New("исправь свой ответ, а лучше ложись поспать")

func currentDayOfTheWeek() string {
	timeNow := TimeNow().Weekday()

	switch {
	case timeNow == time.Monday:
		return "Понедельник"
	case timeNow == time.Tuesday:
		return "Вторник"
	case timeNow == time.Wednesday:
		return "Среда"
	case timeNow == time.Thursday:
		return "Четверг"
	case timeNow == time.Friday:
		return "Пятница"
	case timeNow == time.Saturday:
		return "Суббота"
	case timeNow == time.Sunday:
		return "Воскресенье"
	}

	return "ERROR"
}

func dayOrNight() string {
	timeNow := TimeNow().Hour()

	if int(timeNow) >= 10 && int(timeNow) < 22 {
		return "День"
	} 

	return "Ночь"
}

func nextFriday() int {
	today := currentDayOfTheWeek()
	
	switch {
	case today == "Пятница":
		return 0
	case today == "Четверг":
		return 1
	case today == "Среда":
		return 2
	case today == "Вторник":
		return 3
	case today == "Понедельник":
		return 4
	case today == "Воскресенье":
		return 5
	case today == "Суббота":
		return 6
	}

	return 0
}

func CheckCurrentDayOfTheWeek(answer string) (bool) {
	if utf8.RuneCountInString(answer) > 10 {
		return false
	}

	if answer == "" {
		return false
	}

	answerForCheck := strings.ToLower(currentDayOfTheWeek())

	if strings.ToLower(answer) == answerForCheck {
		return true
	}

	return false
}

func CheckNowDayOrNight(answer string) (bool, error) {
	if utf8.RuneCountInString(answer) > 4 || utf8.RuneCountInString(answer) < 4 {
		return false, IncorrectInput
	}

	answerForCheck := strings.ToLower(answer)
	timeNow := strings.ToLower(dayOrNight())

	if answerForCheck == timeNow {
		return true, nil
	}

	return false, nil

}
