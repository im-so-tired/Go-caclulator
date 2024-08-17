package actions

import (
	"github.com/im-so-tired/Go-caclulator/errors"
	. "github.com/im-so-tired/Go-caclulator/types"
)

type ActionCode = uint

var ActionMap = map[string]ActionCode{
	"+": SumCode,
	"-": SubCode,
	"*": MultiCode,
	"/": DivCode,
}

func sum(a, b Number) Number {
	return a + b
}

func sub(a, b Number) Number {
	return a - b
}

func multi(a, b Number) Number {
	return a * b
}

func division(a, b Number) Number {
	return a / b
}

func GetAction(actionCode ActionCode) (func(a, b Number) Number, error) {
	switch actionCode {
	case SumCode:
		return sum, nil
	case SubCode:
		return sub, nil
	case MultiCode:
		return multi, nil
	case DivCode:
		return division, nil
	default:
		return nil, errors.CalculatorError{Code: errors.SelectedActionNonexistent}
	}
}
