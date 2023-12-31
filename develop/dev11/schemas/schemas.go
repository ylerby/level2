package schemas

type ErrorResponseSchema struct {
	Error string
}

type ResultResponseSchema struct {
	Result interface{}
}

type CreateUpdateDeleteEventRequest struct {
	EventName string `json:"event_name"`
	UserId    int    `json:"user_id"`
	Date      string `json:"date"`
}
