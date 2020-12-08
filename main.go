package main

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

func main() {
	// примерный план работы программы:
	// 1. Считать данные из rule.json
	// 2.
	//
	//
	//

	// Считываем данные из rules.json
	rule, err := readRule("assets/rule.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = rule.FulfillOperandsFields("assets/example.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	functions := map[string]govaluate.ExpressionFunction{
		"in": in,
	}

	evalString, params, err := rule.CreateEvaluationString()
	if err != nil {
		fmt.Printf("failed to create evaluation string: %v\n", err)
	}

	exp, err := govaluate.NewEvaluableExpressionWithFunctions(evalString, functions)
	if err != nil {
		fmt.Println(err)
		return
	}

	var result interface{}
	result, err = exp.Evaluate(params)
	if err != nil {
		fmt.Printf("failed to evaluate string: %v\n", err)
		return
	}

	fmt.Printf("Result is: %v\n", result)
}
