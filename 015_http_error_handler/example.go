package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var ultimateQuestion = "Answer to the Ultimate Question of Life, the Universe, and Everything"

type input struct {
	Question string `json:"question"`
}

type output struct {
	Answer string `json:"answer"`
}

type Oracle struct{}

func (o Oracle) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return NewMethodNotAllowed(r.Method)
	}

	if r.Body == nil {
		return NewBadRequestError("body cannot be nil")
	}
	defer r.Body.Close()

	data := input{}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return NewBadRequestError(err.Error())
	}

	if data.Question == ultimateQuestion {
		response, _ := json.Marshal(output{
			Answer: "42",
		})
		w.Write(response)
	}

	return nil
}

func NewBadRequestError(msg string) error {
	return NewHTTPError(http.StatusBadRequest, msg)
}

func NewMethodNotAllowed(method string) error {
	return NewHTTPError(http.StatusMethodNotAllowed, fmt.Sprintf("method %s not allowed", method))
}
