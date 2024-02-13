package expressions

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrorNotAllowedSymbols = errors.New("expression contains not allowed symbols")
	ErrorWithParentheses   = errors.New("expression contains some eroors with parentheses")
	ErrorRepeatedOperators = errors.New("expression contains repeated operators")
	ErrorEmptyExpression   = errors.New("expression is empty")
	allowedSymbols         = "0123456789+-*/()"
)

func checkForNotAllowedSymbols(expression string) bool {
	for _, symbol := range expression {
		if !strings.Contains(allowedSymbols, string(symbol)) {
			return true
		}
	}
	return false
}
func checkForEmptyExpression(expression string) bool {
	return expression == ""
}

func checkForRepeatedOperators(expression string) bool {
	regex := `[\+\*\/]{2,}|[\+\-\*\/]{2,}[^0-9]`
	matched, _ := regexp.MatchString(regex, expression)
	return matched
}
func checkForParenthesesErrors(expression string) bool {
	regex := `\(\)`
	matched, _ := regexp.MatchString(regex, expression)
	if matched {
		return true
	}
	regex = `\([\+\-\*\/]`
	matched, _ = regexp.MatchString(regex, expression)
	if matched {
		return true
	}
	parentheses := 0
	for _, symbol := range expression {
		if symbol == '(' {
			parentheses++
		}
		if symbol == ')' {
			parentheses--
		}
		if parentheses < 0 {
			return true
		}
	}
	return parentheses != 0
}
func normalizeExpression(expression string) (string, error) {
	expression = strings.ReplaceAll(expression, " ", "")
	if checkForNotAllowedSymbols(expression) {
		return "", ErrorNotAllowedSymbols
	}
	if checkForParenthesesErrors(expression) {
		return "", ErrorWithParentheses
	}
	if checkForRepeatedOperators(expression) {
		return "", ErrorRepeatedOperators
	}
	if checkForEmptyExpression(expression) {
		return "", ErrorEmptyExpression
	}
	return expression, nil
}
