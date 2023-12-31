package responses

import (
	"net/http"
	"os"
)

type ResponseInterface interface {
	JsonResponse(w http.ResponseWriter, httpStatusCode int, data interface{})
}

type Responses struct{}

func (r *Responses) JsonResponse(w http.ResponseWriter, httpStatusCode int, data []byte) {
	w.WriteHeader(httpStatusCode)
	_, err := w.Write(data)
	if err != nil {
		os.Exit(1)
	}
}
