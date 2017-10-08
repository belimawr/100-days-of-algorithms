package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Oracle_ServeHTTP(t *testing.T) {
	expectedAnswer := "42"
	expectedCode := http.StatusOK

	requestBody, _ := json.Marshal(input{
		Question: ultimateQuestion,
	})

	o := Oracle{}

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))

	if err := o.ServeHTTP(w, r); err != nil {
		t.Errorf("did not expect an error: %s", err.Error())
	}

	result := w.Result()
	out := output{}

	if err := json.NewDecoder(result.Body).Decode(&out); err != nil {
		t.Errorf("did not expect an error %s", err.Error())
	}

	if result.StatusCode != expectedCode {
		t.Errorf("expecting status code %d, got %d", expectedCode, result.StatusCode)
	}

	if out.Answer != expectedAnswer {
		t.Errorf("expecting %q, got %q", out.Answer, expectedAnswer)
	}
}

func Test_Oracle_ServeHTTP_MethodNotAllowed(t *testing.T) {
	expectedCode := http.StatusMethodNotAllowed

	o := Oracle{}

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("BREW", "/", nil)

	err := o.ServeHTTP(w, r)
	if err == nil {
		t.Error("expecting an error")
	}

	e, ok := err.(HTTPError)

	if !ok {
		t.Error("expecting an error of type HTTPError")
	}

	if e.Code() != expectedCode {
		t.Errorf("expecting status code %d, got %d", expectedCode, e.Code())
	}
}

func Test_Oracle_ServeHTTP_nilBody(t *testing.T) {
	expectedCode := http.StatusBadRequest

	o := Oracle{}

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/", nil)

	err := o.ServeHTTP(w, r)
	if err == nil {
		t.Error("expecting an error")
	}

	e, ok := err.(HTTPError)

	if !ok {
		t.Error("expecting an error of type HTTPError")
	}

	if e.Code() != expectedCode {
		t.Errorf("expecting status code %d, got %d", expectedCode, e.Code())
	}
}

func Test_Oracle_ServeHTTP_DecodeError(t *testing.T) {
	expectedCode := http.StatusBadRequest

	o := Oracle{}

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/", bytes.NewReader([]byte(`{{ INVALID`)))

	err := o.ServeHTTP(w, r)
	if err == nil {
		t.Error("expecting an error")
	}

	e, ok := err.(HTTPError)

	if !ok {
		t.Error("expecting an error of type HTTPError")
	}

	if e.Code() != expectedCode {
		t.Errorf("expecting status code %d, got %d", expectedCode, e.Code())
	}
}
