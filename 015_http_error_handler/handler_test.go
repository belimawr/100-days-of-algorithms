package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_httpError_Code(t *testing.T) {
	expected := 42
	e := httpError{
		code: expected,
	}

	code := e.Code()

	if code != expected {
		t.Errorf("Expecting %d, got %d", expected, code)
	}
}

func Test_httpError_Error(t *testing.T) {
	expected := "Some Error"
	e := httpError{
		msg: expected,
	}

	msg := e.Error()

	if msg != expected {
		t.Errorf("Expecting %q, got %q", expected, msg)
	}
}

func Test_NewHTTPError(t *testing.T) {
	err := NewHTTPError(http.StatusTeapot, "I'm a teapot")

	if _, ok := err.(httpError); !ok {
		t.Error("Expecting an error of type httpError")
	}
}

func Test_handler_ServeHTTP(t *testing.T) {
	h := HandlerWrapper(Func(func(w http.ResponseWriter, r *http.Request) error {
		return nil
	}))

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("BREW", "/", nil)

	h.ServeHTTP(w, r)

	result := w.Result()

	if result.StatusCode != http.StatusOK {
		t.Errorf("expecting status code 200, got %d", result.StatusCode)
	}
}

func Test_handler_ServeHTTP_error(t *testing.T) {
	expectedCode := http.StatusInternalServerError
	expectedMessage := "Any error message"

	h := HandlerWrapper(Func(func(w http.ResponseWriter, r *http.Request) error {
		return errors.New(expectedMessage)
	}))

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("BREW", "/", nil)

	h.ServeHTTP(w, r)

	result := w.Result()

	body, _ := ioutil.ReadAll(result.Body)

	if result.StatusCode != expectedCode {
		t.Errorf("expecting status code 200, got %d", result.StatusCode)
	}

	if string(body) != expectedMessage {
		t.Errorf("expecting %q, got %q", string(body), expectedMessage)
	}
}

func Test_handler_ServeHTTP_HTTPError(t *testing.T) {
	expectedCode := http.StatusTeapot
	expectedMessage := "I'm a teapot"

	h := HandlerWrapper(Func(func(w http.ResponseWriter, r *http.Request) error {
		return NewHTTPError(expectedCode, expectedMessage)
	}))

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("BREW", "/", nil)

	h.ServeHTTP(w, r)

	result := w.Result()

	body, _ := ioutil.ReadAll(result.Body)

	if result.StatusCode != expectedCode {
		t.Errorf("expecting status code 200, got %d", result.StatusCode)
	}

	if string(body) != expectedMessage {
		t.Errorf("expecting %q, got %q", string(body), expectedMessage)
	}
}
