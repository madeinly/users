package parser

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Parser interface {
	HasErrors() bool
	GetErrors() []Error
	AddError(err string, message string, field string)
	AddErrors(err string, message []string, field string)
	RespondErrors(w http.ResponseWriter)
	Error() string
}

type Error struct {
	Code    string `json:"code"`
	Context string `json:"field,omitempty"`
	Message string `json:"message"`
}

type parser struct {
	errors []Error
}

func New() Parser {
	return &parser{
		errors: make([]Error, 0),
	}
}

func (p *parser) NoError() []Error {
	return make([]Error, 0)
}

func (p *parser) AddError(errorMsg, message string, field string) {
	e := Error{
		Code:    errorMsg,
		Message: message,
		Context: field,
	}

	p.errors = append(p.errors, e)
}

func (p *parser) AddErrors(code string, warnings []string, field string) {

	concatWarning := strings.Join(warnings, ", ")

	p.AddError(code, concatWarning, field)

}

func (p *parser) GetErrors() []Error {
	return p.errors
}

func (p *parser) HasErrors() bool {
	return len(p.errors) > 0
}

func (p *parser) Error() string {

	if p.errors == nil {
		return ""
	}

	errorJson, _ := json.Marshal(p.errors)
	return string(errorJson)
}

func (p *parser) RespondErrors(w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(p.GetErrors())
}
