package app

import (
	"dev11/schemas"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func (a *Application) EventsForDayHandler(w http.ResponseWriter, r *http.Request) {

}

func (a *Application) EventsForWeekHandler(w http.ResponseWriter, r *http.Request) {

}

func (a *Application) EventsForMonthHandler(w http.ResponseWriter, r *http.Request) {

}

func (a *Application) CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	reader, err := io.ReadAll(r.Body)
	if err != nil {
		response, err := json.Marshal(schemas.ErrorResponseSchema{
			ErrorMessage: "reader error",
		})
		if err != nil {
			os.Exit(1)
		}
		a.Response.JsonResponse(w, http.StatusBadRequest, response)
	}

	currentRequestBody := &schemas.CreateEventRequest{}

	err = json.Unmarshal(reader, currentRequestBody)
	if err != nil {
		response, err := json.Marshal(schemas.ErrorResponseSchema{
			ErrorMessage: fmt.Sprintf("unmarshall error - %s", err),
		})
		if err != nil {
			os.Exit(1)
		}
		a.Response.JsonResponse(w, http.StatusInternalServerError, response)
		return
	}

	currentEvent, eventName := a.Validator.CreateEventRequestValidation(currentRequestBody)
	err = a.Calendar.Create(eventName, currentEvent)
	if err != nil {
		response, err := json.Marshal(schemas.ErrorResponseSchema{
			ErrorMessage: fmt.Sprintf("create error - %s", err),
		})
		if err != nil {
			os.Exit(1)
		}
		a.Response.JsonResponse(w, http.StatusInternalServerError, response)
		return
	}
	response, err := json.Marshal(schemas.ResultResponseSchema{
		ResultMessage: "element created",
	})
	if err != nil {
		os.Exit(1)
	}
	a.Response.JsonResponse(w, http.StatusOK, response)
	return
}

func (a *Application) UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	reader, err := io.ReadAll(r.Body)
	if err != nil {
		response, err := json.Marshal(schemas.ErrorResponseSchema{
			ErrorMessage: "reader error",
		})
		if err != nil {
			os.Exit(1)
		}
		a.Response.JsonResponse(w, http.StatusBadRequest, response)
	}

	currentRequestBody := &schemas.UpdateEventRequest{}

	err = json.Unmarshal(reader, currentRequestBody)
	if err != nil {
		response, err := json.Marshal(schemas.ErrorResponseSchema{
			ErrorMessage: fmt.Sprintf("unmarshall error - %s", err),
		})
		if err != nil {
			os.Exit(1)
		}
		a.Response.JsonResponse(w, http.StatusInternalServerError, response)
	}

	currentEvent, eventName := a.Validator.UpdateEventRequestValidation(currentRequestBody)
	err = a.Calendar.Update(eventName, currentEvent)
	if err != nil {
		response, err := json.Marshal(schemas.ErrorResponseSchema{
			ErrorMessage: fmt.Sprintf("create error - %s", err),
		})
		if err != nil {
			os.Exit(1)
		}
		a.Response.JsonResponse(w, http.StatusInternalServerError, response)
	}
}

func (a *Application) DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	reader, err := io.ReadAll(r.Body)
	if err != nil {
		response, err := json.Marshal(schemas.ErrorResponseSchema{
			ErrorMessage: "reader error",
		})
		if err != nil {
			os.Exit(1)
		}
		a.Response.JsonResponse(w, http.StatusBadRequest, response)
	}

	currentRequestBody := &schemas.DeleteEventRequest{}

	err = json.Unmarshal(reader, currentRequestBody)
	if err != nil {
		response, err := json.Marshal(schemas.ErrorResponseSchema{
			ErrorMessage: fmt.Sprintf("unmarshall error - %s", err),
		})
		if err != nil {
			os.Exit(1)
		}
		a.Response.JsonResponse(w, http.StatusInternalServerError, response)
	}

	currentEvent, eventName := a.Validator.DeleteEventRequestValidation(currentRequestBody)
	err = a.Calendar.Delete(eventName, currentEvent)
	if err != nil {
		response, err := json.Marshal(schemas.ErrorResponseSchema{
			ErrorMessage: fmt.Sprintf("create error - %s", err),
		})
		if err != nil {
			os.Exit(1)
		}
		a.Response.JsonResponse(w, http.StatusInternalServerError, response)
	}
}
