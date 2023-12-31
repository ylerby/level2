package schemas

type ErrorResponseSchema struct {
	ErrorMessage string
}

type ResultResponseSchema struct {
	ResultMessage interface{}
}

type CreateEventRequest struct {
	EventName string `json:"event_name"`
	UserId    int    `json:"user_id"`
	Date      string `json:"date"`
}

type UpdateEventRequest struct {
	EventName string `json:"event_name"`
	UserId    int    `json:"user_id"`
	Date      string `json:"date"`
}

type DeleteEventRequest struct {
	EventName string `json:"event_name"`
	UserId    int    `json:"user_id"`
	Date      string `json:"date"`
}
