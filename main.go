package main

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

func main() {
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

	evalString, err := rule.CreateEvaluationString()
	if err != nil {
		fmt.Printf("failed to create evaluation string: %v\n", err)
	}

	exp, err := govaluate.NewEvaluableExpression(evalString)
	if err != nil {
		fmt.Println(err)
		return
	}

	var result interface{}
	result, err = exp.Evaluate(nil)
	if err != nil {
		fmt.Printf("failed to evaluate string: %v\n", err)
		return
	}

	fmt.Printf("Result is: %v\n", result)
}
