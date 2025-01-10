package evaluator

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Convert infix expression to postfix
func InfixToPostfix(expression string) (string, error) {
	stack := []rune{}
	output := []rune{}
	precedence := map[rune]int{
		'+': 1, '-': 1,
		'*': 2, '/': 2, '%': 2,
	}
	isPreviousOperator := true // To handle unary operators at the start or after other operators

	for _, char := range expression {
		if (char >= '0' && char <= '9') || char == '.' { // Operand
			output = append(output, char)
			isPreviousOperator = false
		} else if char == '(' { // Open parenthesis
			stack = append(stack, char)
			isPreviousOperator = true
		} else if char == ')' { // Close parenthesis
			for len(stack) > 0 && stack[len(stack)-1] != '(' {
				output = append(output, ' ')
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				return "", fmt.Errorf("mismatched parentheses")
			}
			stack = stack[:len(stack)-1] // Pop '('
			isPreviousOperator = false
		} else { // Operator
			if isPreviousOperator && char == '-' { // Unary minus
				output = append(output, ' ') // Add space before unary operator
				output = append(output, '~') // Use '~' as a placeholder for unary minus
			} else {
				output = append(output, ' ') // Separate numbers
				for len(stack) > 0 && precedence[stack[len(stack)-1]] >= precedence[char] {
					output = append(output, stack[len(stack)-1])
					output = append(output, ' ')
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, char)
			}
			isPreviousOperator = true
		}
	}

	// Pop remaining operators in the stack
	for len(stack) > 0 {
		if stack[len(stack)-1] == '(' {
			return "", fmt.Errorf("mismatched parentheses")
		}
		output = append(output, ' ')
		output = append(output, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return string(output), nil
}

// Evaluate a postfix expression
func EvaluatePostfix(postfix string) (float64, error) {
	stack := []float64{}
	tokens := strings.Fields(postfix)

	for _, token := range tokens {
		if num, err := strconv.ParseFloat(token, 64); err == nil { // Operand
			stack = append(stack, num)
		} else { // Operator
			if token == "~" { // Unary minus
				if len(stack) < 1 {
					return 0, fmt.Errorf("invalid postfix expression")
				}
				val := stack[len(stack)-1]
				stack[len(stack)-1] = -val
			} else {
				if len(stack) < 2 {
					return 0, fmt.Errorf("invalid postfix expression")
				}
				b := stack[len(stack)-1]
				a := stack[len(stack)-2]
				stack = stack[:len(stack)-2]

				switch token {
				case "+":
					stack = append(stack, a+b)
				case "-":
					stack = append(stack, a-b)
				case "*":
					stack = append(stack, a*b)
				case "/":
					if b == 0 {
						return 0, fmt.Errorf("division by zero")
					}
					stack = append(stack, a/b)
				case "%":
					stack = append(stack, math.Mod(a, b))
				default:
					return 0, fmt.Errorf("unsupported operator: %s", token)
				}
			}
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("invalid postfix expression")
	}

	return stack[0], nil
}

// Evaluate an infix expression
func EvaluateExpression(expression string) (float64, error) {
	postfix, err := InfixToPostfix(expression)
	if err != nil {
		return 0, err
	}
	return EvaluatePostfix(postfix)
}