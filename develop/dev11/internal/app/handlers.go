package app

import (
	"dev11/schemas"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func (a *Application) CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	reader, err := io.ReadAll(r.Body)
	if err != nil {
		response, err := json.Marshal(schemas.ErrorResponseSchema{
			Error: "reader error",
		})
		if err != nil {
			os.Exit(1)
		}
		a.Response.JsonResponse(w, http.StatusInternalServerError, response)
	}

	currentRequestBody := &schemas.CreateUpdateDeleteEventRequest{}

	err = json.Unmarshal(reader, currentRequestBody)
	if err != nil {
		response, err := json.Marshal(schemas.ErrorResponseSchema{
			Error: fmt.Sprintf("unmarshall error - %s", err),
		})
		if err != nil {
			os.Exit(1)
		}
		a.Response.JsonResponse(w, http.StatusInternalServerError, response)
		return
	}

	currentEvent, date, validationErr := a.Validator.CreateUpdateDeleteEventRequestValidation(currentRequestBody)
	err = a.Calendar.Create(currentEvent, date)
	if err != nil || validationErr != nil {
		ErrorMessage := ""
		if err != nil {
			ErrorMessage = fmt.Sprintf("create error - %s", err)
		}
		if validationErr != nil {
			ErrorMessage = fmt.Sprintf("validation error - %s", validationErr)
		}

		response, err := json.Marshal(schemas.ErrorResponseSchema{
			Error: ErrorMessage,
		})
		if err != nil {
			os.Exit(1)
		}
		a.Response.JsonResponse(w, http.StatusServiceUnavailable, response)
		return
	}
	response, err := json.Marshal(schemas.ResultResponseSchema{
		Result: "element created",
	})
	if err != nil {
		os.Exit(1)
	}
	a.Response.JsonResponse(w, http.StatusOK, response)
}

func (a *Application) UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	reader, err := io.ReadAll(r.Body)
	if err != nil {
		response, err := json.Marshal(schemas.ErrorResponseSchema{
			Error: "reader error",
		})
		if err != nil {
			os.Exit(1)
		}
		a.Response.JsonResponse(w, http.StatusBadRequest, response)
	}

	currentRequestBody := &schemas.CreateUpdateDeleteEventRequest{}

	err = json.Unmarshal(reader, currentRequestBody)
	if err != nil {
		response, err := json.Marshal(schemas.ErrorResponseSchema{
			Error: fmt.Sprintf("unmarshall error - %s", err),
		})
		if err != nil {
			os.Exit(1)
		}
		a.Response.JsonResponse(w, http.StatusInternalServerError, response)
	}

	currentEvent, date, validationErr := a.Validator.CreateUpdateDeleteEventRequestValidation(currentRequestBody)
	err = a.Calendar.Update(currentEvent, date)
	if err != nil || validationErr != nil {
		ErrorMessage := ""
		if err != nil {
			ErrorMessage = fmt.Sprintf("create error - %s", err)
		}
		if validationErr != nil {
			ErrorMessage = fmt.Sprintf("validation error - %s", validationErr)
		}

		response, err := json.Marshal(schemas.ErrorResponseSchema{
			Error: ErrorMessage,
		})
		if err != nil {
			os.Exit(1)
		}
		a.Response.JsonResponse(w, http.StatusServiceUnavailable, response)
		return
	}

	response, err := json.Marshal(schemas.ResultResponseSchema{
		Result: "element updated",
	})
	if err != nil {
		os.Exit(1)
	}
	a.Response.JsonResponse(w, http.StatusOK, response)
}

func (a *Application) DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	reader, err := io.ReadAll(r.Body)
	if err != nil {
		response, err := json.Marshal(schemas.ErrorResponseSchema{
			Error: "reader error",
		})
		if err != nil {
			os.Exit(1)
		}
		a.Response.JsonResponse(w, http.StatusBadRequest, response)
	}

	currentRequestBody := &schemas.CreateUpdateDeleteEventRequest{}

	err = json.Unmarshal(reader, currentRequestBody)
	if err != nil {
		response, err := json.Marshal(schemas.ErrorResponseSchema{
			Error: fmt.Sprintf("unmarshall error - %s", err),
		})
		if err != nil {
			os.Exit(1)
		}
		a.Response.JsonResponse(w, http.StatusInternalServerError, response)
	}

	currentEvent, date, validationErr := a.Validator.CreateUpdateDeleteEventRequestValidation(currentRequestBody)
	err = a.Calendar.Delete(currentEvent, date)
	if err != nil || validationErr != nil {
		ErrorMessage := ""
		if err != nil {
			ErrorMessage = fmt.Sprintf("create error - %s", err)
		}
		if validationErr != nil {
			ErrorMessage = fmt.Sprintf("validation error - %s", validationErr)
		}

		response, err := json.Marshal(schemas.ErrorResponseSchema{
			Error: ErrorMessage,
		})
		if err != nil {
			os.Exit(1)
		}
		a.Response.JsonResponse(w, http.StatusServiceUnavailable, response)
		return
	}
	response, err := json.Marshal(schemas.ResultResponseSchema{
		Result: "element deleted",
	})
	if err != nil {
		os.Exit(1)
	}
	a.Response.JsonResponse(w, http.StatusOK, response)
}

func (a *Application) EventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	day := r.URL.Query().Get("day")
	month := r.URL.Query().Get("month")
	year := r.URL.Query().Get("year")

	if day == "" || month == "" || year == "" {
		response, err := json.Marshal(schemas.ErrorResponseSchema{
			Error: "invalid query params",
		})
		if err != nil {
			os.Exit(1)
		}
		a.Response.JsonResponse(w, http.StatusBadRequest, response)
		return
	}

	result, err := a.Calendar.GetEventsForDay(day, month, year)

	if err != nil {
		response, err := json.Marshal(schemas.ErrorResponseSchema{
			Error: fmt.Sprintf("response error - %s", err),
		})
		if err != nil {
			os.Exit(1)
		}
		a.Response.JsonResponse(w, http.StatusInternalServerError, response)
		return
	}

	response, err := json.Marshal(schemas.ResultResponseSchema{
		Result: result,
	})
	if err != nil {
		os.Exit(1)
	}
	a.Response.JsonResponse(w, http.StatusOK, response)
}

func (a *Application) EventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	day := r.URL.Query().Get("day")
	month := r.URL.Query().Get("month")
	year := r.URL.Query().Get("year")

	if day == "" || month == "" || year == "" {
		response, err := json.Marshal(schemas.ErrorResponseSchema{
			Error: "invalid query params",
		})
		if err != nil {
			os.Exit(1)
		}
		a.Response.JsonResponse(w, http.StatusBadRequest, response)
		return
	}

	result, err := a.Calendar.GetEventsForWeek(day, month, year)

	if err != nil {
		response, err := json.Marshal(schemas.ErrorResponseSchema{
			Error: fmt.Sprintf("response error - %s", err),
		})
		if err != nil {
			os.Exit(1)
		}
		a.Response.JsonResponse(w, http.StatusInternalServerError, response)
		return
	}

	response, err := json.Marshal(schemas.ResultResponseSchema{
		Result: result,
	})
	if err != nil {
		os.Exit(1)
	}
	a.Response.JsonResponse(w, http.StatusOK, response)
}

func (a *Application) EventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	month := r.URL.Query().Get("month")
	year := r.URL.Query().Get("year")

	if month == "" || year == "" {
		response, err := json.Marshal(schemas.ErrorResponseSchema{
			Error: "invalid query params",
		})
		if err != nil {
			os.Exit(1)
		}
		a.Response.JsonResponse(w, http.StatusBadRequest, response)
		return
	}

	result, err := a.Calendar.GetEventsForMonth(month, year)

	if err != nil {
		response, err := json.Marshal(schemas.ErrorResponseSchema{
			Error: fmt.Sprintf("response error - %s", err),
		})
		if err != nil {
			os.Exit(1)
		}
		a.Response.JsonResponse(w, http.StatusInternalServerError, response)
		return
	}

	response, err := json.Marshal(schemas.ResultResponseSchema{
		Result: result,
	})
	if err != nil {
		os.Exit(1)
	}
	a.Response.JsonResponse(w, http.StatusOK, response)
}
