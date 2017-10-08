package main

import "net/http"

// ErrorHandler an interface for HTTP handlers that returns an error
type ErrorHandler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request) error
}

// HandlerWrapper wrapper to convert an ErrorHandler in http.Handler
func HandlerWrapper(h ErrorHandler) http.Handler {
	return handler{
		h: h,
	}
}

// Func a wrapper to use a function as a ErrorHandler
type Func func(w http.ResponseWriter, r *http.Request) error

func (f Func) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	return f(w, r)
}

type handler struct {
	h ErrorHandler
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.h.ServeHTTP(w, r)

	if err != nil {
		switch e := err.(type) {
		case HTTPError:
			w.WriteHeader(e.Code())
			w.Write([]byte(e.Error()))
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}

// HTTPError an interface that satisfies the error interface
// and has a method that returns the appropriated HTTP status code
type HTTPError interface {
	error
	Code() int
}

type httpError struct {
	code int
	msg  string
}

func (e httpError) Code() int {
	return e.code
}

func (e httpError) Error() string {
	return e.msg
}

// NewHTTPError returns an error that satisfies HTTPError interface
func NewHTTPError(code int, msg string) HTTPError {
	return httpError{
		code: code,
		msg:  msg,
	}
}
