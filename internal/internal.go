package internal

import (
	"strings"

	"github.com/overseven/go-math-expression-parser/interfaces"
)

func PrepareString(str string) string {
	str = strings.ReplaceAll(str, " ", "")
	str = strings.TrimSpace(str)
	return str
}

func unaryOperatorExist(op string, p interfaces.ExpParser) (index int, exist bool) {
	if _, ok := p.GetFunctions()[0][op]; ok {
		return 0, true
	}
	return -1, false
}

func binaryOperatorExist(op string, p interfaces.ExpParser) (index int, exist bool) {
	for i := 1; i <= 2; i++ {
		if _, ok := p.GetFunctions()[i][op]; ok {
			return i, true
		}
	}
	return -1, false
}

// ParenthesisIsCorrect - checks correct parenthesis pairs
func ParenthesisIsCorrect(str string) (index int, correct bool) {
	counter := 0
	for i, c := range str {
		switch c {
		case '(':
			counter++
		case ')':
			counter--
		default:
			continue
		}
		if counter < 0 {
			return i, false
		}
	}
	if counter > 0 {
		return len(str) - 1, false
	}
	return -1, true
}
