package main

import (
	validate "mathOp/cmd"
	"testing"
)

// Addition validation argument
func TestAdditionArgumentValidationWhenBothArgumentsAreNumbers(t *testing.T) {
	args := []string{"1", "2"}
	isValid, _ := validate.ValidateAdditionArgs(args)
	if !isValid {
		t.Errorf("Addition Validation of args: 1,2 is incorrect, got false but should be true")
	}
}

func TestAdditionArgumentValidation_MoreThanTwoNumbers(t *testing.T) {
	args := []string{"1", "2", "3"}
	isValid, _ := validate.ValidateAdditionArgs(args)
	if isValid {
		t.Errorf("Addition Validation of args: 1,2,3 is incorrect, got TRUE but should be FALSE")
	}
}

func TestAdditionArgumentValidation_notNumbers(t *testing.T) {
	args := []string{"1", "hbjsdf"}
	isValid, _ := validate.ValidateAdditionArgs(args)
	if isValid {
		t.Errorf("Addition Validation of args: 1,hbjsdf, is incorrect, got TRUE but should be FALSE")
	}
}

// Sqrt validation argument
func TestSqrtArgumentValidation_ArgumentIsANumber(t *testing.T) {
	args := []string{"64"}
	isValid, _ := validate.ValidateSqrtArgs(args)
	if !isValid {
		t.Errorf("Sqrt Validation of args: 8 is incorrect, got false but should be true")
	}
}

func TestSqrtArgumentValidation_MoreThanOneNumbers(t *testing.T) {
	args := []string{"1", "2", "3"}
	isValid, _ := validate.ValidateSqrtArgs(args)
	if isValid {
		t.Errorf("Sqrt Validation of args: 1,2,3 is incorrect, got TRUE but should be FALSE")
	}
}

func TestSqrtArgumentValidation_notNumbers(t *testing.T) {
	args := []string{"HDJFS"}
	isValid, _ := validate.ValidateSqrtArgs(args)
	if isValid {
		t.Errorf("Sqrt Validation of args: HDJFS, is incorrect, got TRUE but should be FALSE")
	}
}
