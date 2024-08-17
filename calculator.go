package main

import (
	"github.com/im-so-tired/Go-caclulator/actions"
	errors "github.com/im-so-tired/Go-caclulator/errors"
	. "github.com/im-so-tired/Go-caclulator/types"
)

type Calculator interface {
	inputValue(value Number) error
	chooseActions(action string) error
	getResult() Number
	//cancelLastAction()
}

type BaseCalculatorV2 struct {
	isInitial    bool
	result       Number
	action       uint
	enteredValue *Number
}

func (c *BaseCalculatorV2) inputValue(value Number) error {
	var err error

	if c.isInitial {
		c.result = value
		c.isInitial = false
		return nil
	}

	if c.action != 0 {
		c.enteredValue = &value
		err = c.Calculate()
	} else {
		err = errors.CalculatorError{Code: errors.ValueAlreadyEntered}
	}

	return err
}

func (c *BaseCalculatorV2) chooseActions(action string) error {
	if actionCode, ok := actions.ActionMap[action]; ok {
		c.action = actionCode
	} else {
		return errors.CalculatorError{Code: errors.SelectedActionNonexistent}
	}

	return nil
}

func (c *BaseCalculatorV2) getResult() Number {
	return c.result
}

func (c *BaseCalculatorV2) Calculate() error {
	if c.action == 0 || c.enteredValue == nil {
		return nil
	}

	actionMethod, err := actions.GetAction(c.action)

	if err != nil {
		return err
	}

	c.result = actionMethod(c.result, *c.enteredValue)

	c.action = 0
	c.enteredValue = nil

	return nil
}

var CalculatorInstance Calculator = &BaseCalculatorV2{isInitial: true}
