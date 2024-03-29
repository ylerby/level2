package app

import (
	"dev11/internal/middleware"
	"dev11/internal/responses"
	"dev11/internal/service/calendar"
	"dev11/internal/service/validation"
	"net/http"
	"os"
)

type Application struct {
	Server    *http.Server
	Calendar  *calendar.Calendar
	Response  *responses.Responses
	Validator *validation.Validator
}

func New(port string) *Application {
	return &Application{
		Server: &http.Server{
			Addr: port,
		},
		Calendar: &calendar.Calendar{
			CalendarMap: make(map[string][]calendar.Events),
		},
		Response:  &responses.Responses{},
		Validator: &validation.Validator{},
	}
}

func (a *Application) Run() {

	http.HandleFunc("/events_for_day", middleware.LoggingMiddleware("GET", a.EventsForDayHandler))
	http.HandleFunc("/events_for_week", middleware.LoggingMiddleware("GET", a.EventsForWeekHandler))
	http.HandleFunc("/events_for_month", middleware.LoggingMiddleware("GET", a.EventsForMonthHandler))
	http.HandleFunc("/create_events", middleware.LoggingMiddleware("POST", a.CreateEventHandler))
	http.HandleFunc("/update_events", middleware.LoggingMiddleware("POST", a.UpdateEventHandler))
	http.HandleFunc("/delete_events", middleware.LoggingMiddleware("POST", a.DeleteEventHandler))

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		os.Exit(1)
	}
}
