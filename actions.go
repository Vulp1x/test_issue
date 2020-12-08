package main

import "fmt"

// in function checks if the provided value is stored in the provided list
// the value and list should be of the same type e.g
// value: string, list: []string
func in(args ...interface{}) (interface{}, error) {
	if argsLength := len(args); argsLength != 2 {
		return nil, fmt.Errorf("expected 2 arguments, but got: %d", argsLength)
	}

	switch castedValue := args[0].(type) {
	case int:
		castedList, ok := args[1].([]int)
		if !ok {
			return nil, fmt.Errorf("expected []int type for list parameter, but got %T type", args[1])
		}

		for _, value := range castedList {
			if value == castedValue {
				return true, nil
			}
		}

	case string:
		castedList, ok := args[1].([]string)
		if !ok {
			return nil, fmt.Errorf("expected []string type for list parameter, but got %T type", args[1])
		}

		for _, value := range castedList {
			if value == castedValue {
				return true, nil
			}
		}
	default:
		return nil, fmt.Errorf("expected int or string type for a value parameter, but got %T type", args[0])
	}

	return false, nil
}
