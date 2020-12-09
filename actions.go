package main

import (
	"fmt"
	"strings"
)

// in function checks if the provided value is stored in the provided list
// the value and list should be of the same type e.g
// value: string, list: []string
func in(value interface{}, slice []interface{}) string {
	buf := strings.Builder{}
	buf.WriteString("(")

	for i, sliceValue := range slice {
		buf.WriteString(fmt.Sprintf("('%v' == '%v')", value, sliceValue))
		if i < len(slice)-1 {
			buf.WriteString(" || ")
		}
	}

	buf.WriteString(")")

	return buf.String()
}
