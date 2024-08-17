package errors

const (
	ValueAlreadyEntered ErrorCode = iota + 1
	DivisionByZero
	SelectedActionNonexistent
)

var MessagesMap = map[ErrorCode]string{
	ValueAlreadyEntered:       "Значение уже введено",
	DivisionByZero:            "Невозможно деление на 0",
	SelectedActionNonexistent: "Выбрано несуществующее действие",
}
