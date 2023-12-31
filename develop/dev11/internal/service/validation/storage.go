package validation

import (
	"dev11/internal/service/calendar"
	"dev11/schemas"
)

type ValidatorServiceInterface interface {
	CreateUpdateDeleteEventRequestValidation(r *schemas.CreateUpdateDeleteEventRequest) (calendar.Events, string, error)
	GetRequestValidation(day, month string) (string, string, error)
}

type Validator struct{}
