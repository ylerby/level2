package validation

import (
	"dev11/internal/service/calendar"
	"dev11/schemas"
	"fmt"
	"strconv"
	"strings"
)

func (v *Validator) CreateUpdateDeleteEventRequestValidation(r *schemas.CreateUpdateDeleteEventRequest) (calendar.Events, string, error) {
	splitDate := strings.Split(r.Date, "-")

	convertDay, dayConvertError := strconv.Atoi(splitDate[0])
	convertMonth, monthConvertError := strconv.Atoi(splitDate[1])
	convertYear, yearConvertError := strconv.Atoi(splitDate[2])

	if dayConvertError != nil || monthConvertError != nil || yearConvertError != nil {
		return calendar.Events{}, "", fmt.Errorf("invalid data")
	}

	if convertDay < 1 || convertDay > 31 {
		return calendar.Events{}, "", fmt.Errorf("invalid data")
	}

	if convertMonth < 1 || convertMonth > 12 {
		return calendar.Events{}, "", fmt.Errorf("invalid data")
	}

	if convertYear < 1 {
		return calendar.Events{}, "", fmt.Errorf("invalid data")
	}

	day := ""
	month := ""

	if convertDay < 10 {
		day = fmt.Sprintf("0%d", convertDay)
	} else {
		day = splitDate[0]
	}

	if convertMonth < 10 {
		month = fmt.Sprintf("0%d", convertMonth)
	} else {
		month = splitDate[1]
	}

	date := fmt.Sprintf("%s-%s-%s", day, month, splitDate[2])
	return calendar.Events{
		EventName: r.EventName,
		UserId:    r.UserId,
	}, date, nil
}

func (v *Validator) GetRequestValidation(day, month string) (string, string, error) {
	convertDay, err := strconv.Atoi(day)
	if err != nil {
		return "", "", fmt.Errorf("invalid data")
	}

	convertMonth, err := strconv.Atoi(month)
	if err != nil {
		return "", "", fmt.Errorf("invalid data")
	}

	if convertDay < 10 {
		day = fmt.Sprintf("0%d", convertDay)
	}

	if convertMonth < 10 {
		month = fmt.Sprintf("0%d", convertMonth)
	}

	return day, month, nil
}
