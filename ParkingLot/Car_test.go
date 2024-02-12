package main

import (
	"testing"
)

func TestCreateNewCarWithValidRegistrationNumber(t *testing.T) {
	registrationNumber := "ABC123"
	color := "Red"
	car, error := NewCar(RegistrationNumber(registrationNumber), color)

	if error != nil {
		t.Errorf("Expected no error, but got: %v", error)
	}

	if car == nil {
		t.Error("Expected non-nil Car instance, got nil")
	} else {
		if car.parkingTicket.registrationNumber != RegistrationNumber(registrationNumber) {
			t.Errorf("Expected registration number %s, got %s", registrationNumber, car.parkingTicket.registrationNumber)
		}
		if car.color != color {
			t.Errorf("Expected color %s, got %s", color, car.color)
		}
	}
}

func TestCreateNewCarWithInvalidRegistrationNumber(t *testing.T) {
	registrationNumber := "invalid"
	color := "Red"
	expectedErrorMessage := "Invalid registration number format"

	car, error := NewCar(RegistrationNumber(registrationNumber), color)

	if error == nil {
		t.Error("Expected an error, but got nil")
	} else if error.Error() != expectedErrorMessage {
		t.Errorf("Expected error message '%s', got '%s'", expectedErrorMessage, error.Error())
	}

	if car != nil {
		t.Errorf("Expected nil Car instance, got non-nil")
	}
}
