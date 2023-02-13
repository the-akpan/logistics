package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Body    interface{} `json:"data,omitempty"`
}

func CreateResponse(res http.ResponseWriter) *Response {
	res.Header().Set("Content-Type", "application/json")
	response := Response{}

	return &response
}

func DecodeRequest(form interface{}, response *Response, res http.ResponseWriter, req *http.Request) bool {
	err := json.NewDecoder(req.Body).Decode(form)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		response.Message = "Error Decoding values " + err.Error()
		return false
	}
	return true
}
