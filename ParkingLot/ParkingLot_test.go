package main

import (
	"fmt"
	"testing"
)

func TestCreateNewParkingLotWithInvalidCapacity(t *testing.T) {
	_, error := NewParkingLot(-2)
	if error == nil || error.Error() != "Invalid capacity. Capacity must be a positive value." {
		t.Errorf("Expected Invalid capacity error, got: %v", error)
	}

}

func TestParkCarSuccessfully(t *testing.T) {
	lot, error := NewParkingLot(2)
	if error != nil {
		t.Errorf("Unexpected error: %v", error)
	}

	for i := 0; i < 2; i++ {
		car := Car{
			parkingTicket: ParkingTicket{registrationNumber: RegistrationNumber(fmt.Sprintf("ABC%d", i+100))},
			color:         "Blue",
		}

		err := lot.ParkCar(car)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}

func TestParkCarWhenLotFullReturnsParkingLotFullError(t *testing.T) {
	lot, error := NewParkingLot(2)
	if error != nil {
		t.Errorf("Unexpected error: %v", error)
	}

	for i := 0; i < 2; i++ {
		car := Car{
			parkingTicket: ParkingTicket{registrationNumber: RegistrationNumber(fmt.Sprintf("ABC%d", i+100))},
			color:         "Blue",
		}

		err := lot.ParkCar(car)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	}

	err := lot.ParkCar(Car{
		parkingTicket: ParkingTicket{registrationNumber: "XYZ123"},
		color:         "Red",
	})
	if err == nil || err.Error() != "Parking lot is full. Trying the next one." {
		t.Errorf("Expected parking lot full error, got: %v", err)
	}
}

func TestParkCarWhenCarAlreadyParkedReturnsCarAlreadyParkedError(t *testing.T) {
	lot, error := NewParkingLot(1)
	if error != nil {
		t.Errorf("Unexpected error: %v", error)
	}

	car := Car{
		parkingTicket: ParkingTicket{registrationNumber: "XYZ123"},
		color:         "Blue",
	}
	err := lot.ParkCar(car)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = lot.ParkCar(car)
	if err == nil || err.Error() != "Car with registration number is already parked" {
		t.Errorf("Expected parking lot full error, got: %v", err)
	}
}

func TestUnparkCarSuccessfully(t *testing.T) {
	lot, error := NewParkingLot(2)
	if error != nil {
		t.Errorf("Unexpected error: %v", error)
	}

	car := Car{
		parkingTicket: ParkingTicket{registrationNumber: "ABC123"},
		color:         "Blue",
	}
	err := lot.ParkCar(car)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	_, err = lot.UnparkCar(car.parkingTicket)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestUnparkCarWhenCarNotPresentInLotReturnsCarNotFoundError(t *testing.T) {
	lot, error := NewParkingLot(2)
	if error != nil {
		t.Errorf("Unexpected error: %v", error)
	}

	car := Car{
		parkingTicket: ParkingTicket{registrationNumber: "ABC123"},
		color:         "Blue",
	}
	err := lot.ParkCar(car)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	_, err = lot.UnparkCar(car.parkingTicket)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	_, err = lot.UnparkCar(car.parkingTicket)
	if err == nil || err.Error() != "Car with registration number ABC123 not found" {
		t.Errorf("Expected car not found error, got: %v", err)
	}
}
