package validation

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

var validate = validator.New()

type Field struct {
	Name    string `json:"name,omitempty"`
	Accepts string `json:"accepts,omitempty"`
	Tag     string `json:"tag,omitempty"`
}

type Error struct {
	ValidationError string  `json:"validationError"`
	Fields          []Field `json:"fields,omitempty"`
	Message         string  `json:"message,omitempty"`
	details         error
}

// Struct accepts something and returns a validation err
func Struct(s interface{}, msgAndArgs ...interface{}) *Error {
	var validErr Error
	err := validate.Struct(s)
	if err != nil {
		validErr.ValidationError = err.Error()
		if _, ok := err.(validator.ValidationErrors); !ok {
			validErr = Error{ValidationError: "Cannot validate this object because it is not a struct", Message: messageFromMsgAndArgs(msgAndArgs)}
			return &validErr
		}
		for _, err := range err.(validator.ValidationErrors) {
			validErr.Fields = append(validErr.Fields, Field{Name: err.Field(), Accepts: err.Type().String(), Tag: err.Tag()})
		}
		return &validErr
	}
	return nil
}

func messageFromMsgAndArgs(msgAndArgs ...interface{}) string {
	if len(msgAndArgs) == 0 || msgAndArgs == nil {
		return ""
	}
	if len(msgAndArgs) == 1 {
		msg := msgAndArgs[0]
		if msgAsStr, ok := msg.(string); ok {
			return msgAsStr
		}
		return fmt.Sprintf("%+v", msg)
	}
	if len(msgAndArgs) > 1 {
		return fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
	}
	return ""
}
