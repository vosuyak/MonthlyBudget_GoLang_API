package common

import (
	"encoding/json"
	"net/http"
)

type (
	ServiceError struct {
		Server  string `json:"server"`
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	ResourceError struct {
		Error ServiceError `json:"error"`
	}

	ServiceSuccess struct {
		DB       interface{} `json:"db"`
		Code     int         `json:"code"`
		Message  string      `json:"message"`
		Data interface{} `json:"data"`
	}
	ResourceSuccess struct {
		Response ServiceSuccess `json:"response"`
	}
)

// DisplayError - Pass the writer for the Response, handlerError for database response, pass http.Status code, message for personalize error response
func DisplayError(w http.ResponseWriter, handlerError error, code int, message string) {
	errObject := ServiceError{
		Server:  handlerError.Error(),
		Code:    code,
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if j, err := json.Marshal(ResourceError{Error: errObject}); err == nil {
		w.Write(j)
	}
}

// DisplaySuccess - Pass the writer for the Response, db for C,U,D for database response or true for R, pass http.Status code, Response for R
func DisplaySuccess(w http.ResponseWriter, db interface{}, code int, data interface{}) {
	successObject := ServiceSuccess{
		Code:     code,
		Message:  "success",
		DB:       db,
		Data: data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if j, err := json.Marshal(ResourceSuccess{Response: successObject}); err == nil {
		w.Write(j)
	}
}

// DisplayCustomResponse - Pass the writer for the Response, msg Custom Message, pass http.Status code, Response for R
func DisplayCustomResponse(w http.ResponseWriter, msg string, code int, data interface{}) {
	successObject := ServiceSuccess{
		Code:     code,
		Message:  msg,
		DB:       nil,
		Data: data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if j, err := json.Marshal(ResourceSuccess{Response: successObject}); err == nil {
		w.Write(j)
	}
}
