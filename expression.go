package main

import (
	"fmt"
	"strings"

	"github.com/oliveagle/jsonpath"
)

type Expression struct {
	Action         string
	Left, Right    *Operand
	Operator       Operator
	IsGroupEnabled bool `json:"group"`
}

type Type string

type Operand struct {
	Type      Type `json:"dtype"`
	Value     interface{}
	FieldName string `json:"field"`
}

func (t Type) IsSliceType() bool {
	return strings.Contains(string(t), "list")
}

func (t Type) IsInt() bool {
	return strings.Contains(string(t), "int")
}

func (t Type) IsString() bool {
	return strings.Contains(string(t), "string")
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

func (e Expression) toString() (string, error) {
	err := e.checkOperandsTypes()
	if err != nil {
		return "", err
	}

	if e.Action == "in" {
		return fmt.Sprintf("%s %s ", in(e.Left.Value, e.Right.Value.([]interface{})), e.Operator), nil
	}

	return fmt.Sprintf("(%v %s %v) %s ", e.Left.Value, e.Action, e.Right.Value, e.Operator), nil
}

func (e Expression) checkOperandsTypes() error {
	// string and int couldn't be compared
	if e.Left.Type.IsString() && e.Right.Type.IsInt() ||
		e.Left.Type.IsInt() && e.Right.Type.IsString() {
		return fmt.Errorf("bad operand's types provided: %s and %s", e.Left.Type, e.Right.Type)
	}

	// check that action suits provided types.
	valid := isActionValid(e.Action, e.Left.Type, e.Right.Type)
	if !valid {
		return fmt.Errorf("action %s couldn't be performed on types: %s and %s", e.Action, e.Left.Type, e.Right.Type)
	}

	return nil
}

// is actionValid checks if action could be performed on the specified types.
// leftOperandType and rightOperandType are types of expression operands.
// Assume, that they both contain string or both contain int.
func isActionValid(action string, leftOperandType, rightOperandType Type) bool {
	// left operand has a slice type. There are no available actions for this case.
	if leftOperandType.IsSliceType() {
		return false
	}

	if rightOperandType.IsSliceType() {
		return action == "in"
	}

	switch leftOperandType {
	case "int":
		return action == "<" || action == "<=" || action == "==" || action == ">=" || action == ">"
	case "string":
		return action == "=="
	default:
		// unknown type
		return false
	}
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
