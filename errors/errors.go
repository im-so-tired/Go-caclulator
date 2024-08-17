package errors

type ErrorCode = uint

type CalculatorError struct {
	Code ErrorCode
}

func (e CalculatorError) Error() string {
	message, ok := MessagesMap[e.Code]

	if ok {
		return message
	}

	return "unknown error"
}

func (e CalculatorError) GetCode() ErrorCode {
	return e.Code
}
