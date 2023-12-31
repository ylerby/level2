package calendar

type EventsServiceInterface interface {
	Create(currentEventName string, event Events) error
	Update(currentEventName string, event Events) error
	Delete(currentEventName string, event Events) error
	GetEventsForDay(day string)
	GetEventsForWeek(week string)
	GetEventsForMonth(month string)
}

type Calendar struct {
	CalendarMap map[string]Events
}

type Events struct {
	UserId int
	Date   string
}
