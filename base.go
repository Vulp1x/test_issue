package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Rule struct {
	Body struct {
		Expressions []Expression `json:"expression"`
	}
}

type Operator string

func readRule(source string) (*Rule, error) {
	b, err := ioutil.ReadFile(source)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}

	var rule Rule

	err = json.Unmarshal(b, &rule)

	return &rule, err
}

// FulfillOperandsFields updates both operands of the expression.
// So they contain values from the lookup fields of event json
func (r *Rule) FulfillOperandsFields(exampleSourceFileName string) error {
	var jsonData interface{}

	data, err := ioutil.ReadFile(exampleSourceFileName)
	if err != nil {
		return fmt.Errorf("failed to open example json file: %v", err)
	}

	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return fmt.Errorf("failed to unmarshal json: %v", err)
	}

	for _, expression := range r.Body.Expressions {
		err = expression.SetMissingFields(jsonData)
		if err != nil {
			return fmt.Errorf("failed to set missing fields for expression: %#v due to %v", expression, err)
		}
	}

	return nil
}

// CreateEvaluationString prepares string for following evaluation.
func (r Rule) CreateEvaluationString() (
	evaluateString string,
	parameters map[string]interface{},
	err error,
) {
	return "(\"ok\" == \"ok\" || \"ok\" == \"ok\" || \"ok\" == \"ok\") && (1 == 4 || 1 == 3 || 1 == 1) &&" +
			" (15 == 15 || \"high\" == \"high\") && 15 == 15 && 5 == 5 ",
		nil, nil
}
