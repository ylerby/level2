package validation

import (
	"dev11/internal/service/calendar"
	"dev11/schemas"
)

func (v *Validator) CreateEventRequestValidation(r *schemas.CreateEventRequest) (calendar.Events, string) {
	return calendar.Events{
		Date:   r.Date,
		UserId: r.UserId,
	}, r.EventName
}

func (v *Validator) UpdateEventRequestValidation(r *schemas.UpdateEventRequest) (calendar.Events, string) {
	return calendar.Events{}, ""
}

func (v *Validator) DeleteEventRequestValidation(r *schemas.DeleteEventRequest) (calendar.Events, string) {
	return calendar.Events{}, ""
}
