package responses

import "net/http"

type ResponseInterface interface {
	JsonResponse(w http.ResponseWriter, httpStatusCode int, data interface{})
}

type Responses struct{}

func (r *Responses) JsonResponse(w http.ResponseWriter, httpStatusCode int, data []byte) {
	w.WriteHeader(httpStatusCode)
	w.Write(data)
}
