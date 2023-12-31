package validation

import (
	"dev11/internal/service/calendar"
	"dev11/schemas"
)

type ValidatorServiceInterface interface {
	CreateEventRequestValidation(r *schemas.CreateEventRequest) (calendar.Events, string)
	UpdateEventRequestValidation(r *schemas.UpdateEventRequest) (calendar.Events, string)
	DeleteEventRequestValidation(r *schemas.DeleteEventRequest) (calendar.Events, string)
}

type Validator struct{}
