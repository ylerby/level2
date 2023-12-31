package calendar

type EventsServiceInterface interface {
	Create(event Events, date string) error
	Update(event Events, date string) error
	Delete(event Events, date string) error
	GetEventsForDay(day, month, year string) ([]Events, error)
	GetEventsForWeek(day, month, year string) ([][]Events, error)
	GetEventsForMonth(month, year string) ([][]Events, error)
}

type Calendar struct {
	CalendarMap map[string][]Events
}

type Events struct {
	UserId    int
	EventName string
}
