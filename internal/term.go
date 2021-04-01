package internal

import (
	"errors"
	"strconv"

	"github.com/overseven/go-math-expression-parser/interfaces"
)

// Term - the struct which contains a single value
type Term struct {
	Val string
}

func (t *Term) GetVarList(vars map[string]interface{}) {
	if t.Val == "" {
		return
	}
	if _, err := strconv.ParseFloat(t.Val, 64); err == nil {
		return
	}
	vars[t.Val] = struct{}{}

}

// Evaluate - return a value which contains in Term
func (t *Term) Evaluate(vars map[string]float64, p interfaces.ExpParser) (float64, error) {
	if t.Val == "" {
		return 0.0, nil
	}
	if val, err := strconv.ParseFloat(t.Val, 64); err == nil {
		return val, nil
	}
	val, ok := vars[t.Val]
	if !ok {
		return 0.0, errors.New("value '" + t.Val + " not found in map")
	}
	return val, nil
}

// toString conversation
func (t *Term) String() string {
	return t.Val
}
