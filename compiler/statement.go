// Pseudo statement compiler
package compiler

import "strings"

type PrepareResult int8

const (
	SUCCESS PrepareResult = iota
	UNRECOGNIZED_STATEMENT
)

type Type int

const (
	Insert Type = iota
	Select
	Error
)

type Statement struct {
	Type Type
}

func PrepareStatement(input string) (Statement, PrepareResult) {
	if strings.HasPrefix(input, "insert") {
		return Statement{Insert}, SUCCESS
	} else if strings.HasPrefix(input, "select") {
		return Statement{Select}, SUCCESS
	} else {
		return Statement{Error}, UNRECOGNIZED_STATEMENT
	}
}
