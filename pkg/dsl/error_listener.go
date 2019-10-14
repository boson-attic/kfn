package dsl

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type ErrorListener struct {
	*antlr.DefaultErrorListener
	errors []error
}

func NewErrorListener() *ErrorListener {
	return &ErrorListener{errors: make([]error, 0)}
}

func (e2 *ErrorListener) GetErrors() []error {
	return e2.errors
}

func (e2 *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	e2.errors = append(e2.errors, fmt.Errorf("at %d:%d: %s", line, column, e.GetMessage()))
}
