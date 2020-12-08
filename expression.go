package main

import (
	"strings"

	"github.com/oliveagle/jsonpath"
)

type Expression struct {
	Action         string
	Left, Right    *Operand
	Operator       Operator
	IsGroupEnabled bool `json:"group"`
}

type Operand struct {
	Type      string `json:"dtype"`
	Value     interface{}
	FieldName string `json:"field"`
}

// SetMissingFields lookups jsonData structure to get needed values.
func (e *Expression) SetMissingFields(jsonData interface{}) error {
	if e.Left.FieldName != "" {
		err := e.Left.setMissingField(jsonData)
		if err != nil {
			return err
		}
	}

	if e.Right.FieldName != "" {
		err := e.Right.setMissingField(jsonData)
		if err != nil {
			return err
		}
	}

	return nil
}

func (o *Operand) setMissingField(jsonData interface{}) error {
	o.prepareFieldNameForLookup()

	res, err := jsonpath.JsonPathLookup(jsonData, o.FieldName)
	if err != nil {
		return err
	}

	o.Value = res

	return nil
}

// prepareFieldNameForLookup sets Operand's FieldName for the format of the jsonpath lookup library.
func (o *Operand) prepareFieldNameForLookup() {
	o.FieldName = "$." + strings.ReplaceAll(o.FieldName, "#", "[:]")
}
