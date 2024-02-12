package main

import (
	"testing"
)

func TestParkCarSuccessfully_WhenCurrentLotIsEmpty_OneAttendantAssignedToMultipleLots(t *testing.T) {
	attendant := NewParkingLotAttendant()
	parkingLot1, _ := NewParkingLot(5)
	parkingLot2, _ := NewParkingLot(5)
	attendant.assign(parkingLot1)
	attendant.assign(parkingLot2)

	car, err := NewCar("ABC123", "Red")
	if err != nil {
		t.Fatalf("Error creating car: %v", err)
	}

	err = attendant.ParkCar(*car)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if parkingLot1.isCarParked(*car) {
		t.Logf("Car parked successfully")
	} else {
		t.Error("Car not parked in any assigned lot")
	}
}

// func TestParkCarSuccessfully_WhenCurrentLotIsEmpty_MultipleAttendantsAssignedToOneLot(t *testing.T) {
// 	attendant1 := NewParkingLotAttendant()
// 	attendant2 := NewParkingLotAttendant()
// 	parkingLot, _ := NewParkingLot(5)
// 	attendant1.AddAssignedLot(parkingLot)
// 	attendant2.AddAssignedLot(parkingLot)

// 	car, err := NewCar("ABC123", "Red")
// 	if err != nil {
// 		t.Fatalf("Error creating car: %v", err)
// 	}

// }

func TestParkCarSuccessfully_WhenCurrentLotIsFull_TriesParkingInNextLot(t *testing.T) {
	attendant := NewParkingLotAttendant()
	parkingLot1, _ := NewParkingLot(1)
	parkingLot2, _ := NewParkingLot(2)
	attendant.assign(parkingLot1)
	attendant.assign(parkingLot2)

	car1, err1 := NewCar("ABC123", "Red")
	if err1 != nil {
		t.Fatalf("Error creating car: %v", err1)
	}
	parkingLot1.ParkCar(*car1)

	newCar, err2 := NewCar("XYZ789", "Blue")
	if err2 != nil {
		t.Fatalf("Error creating car: %v", err2)
	}

	err3 := attendant.ParkCar(*newCar)

	if err3 == nil && parkingLot2.isCarParked(*newCar) {
		t.Logf("Car parked successfully in the next available lot")
	} else {
		t.Errorf("Expected car to be parked in the next available lot, got error: %v", err3)
	}
}

func TestParkCar_WhenAllParkingLotsFull_ReturnsAllParkingLotsFullError(t *testing.T) {
	attendant := NewParkingLotAttendant()
	parkingLot1, _ := NewParkingLot(1)
	parkingLot2, _ := NewParkingLot(1)
	attendant.assign(parkingLot1)
	attendant.assign(parkingLot2)

	car1, _ := NewCar("ABC123", "Red")
	car2, _ := NewCar("XYZ789", "Blue")
	parkingLot1.ParkCar(*car1)
	parkingLot2.ParkCar(*car2)

	newCar, err := NewCar("ABC456", "Green")
	if err != nil {
		t.Fatalf("Error creating car: %v", err)
	}

	err = attendant.ParkCar(*newCar)

	expectedErrorMsg := "All parking lots are full. Cannot park the car."
	if err != nil && err.Error() == expectedErrorMsg {
		t.Logf("Expected error received: %v", err)
	} else {
		t.Errorf("Expected error message '%s', got '%v'", expectedErrorMsg, err)
	}
}

func TestAddAssignedLotToParkingLotAttendantSuccessFully(t *testing.T) {
	attendant := NewParkingLotAttendant()
	parkingLot, _ := NewParkingLot(5)

	attendant.assign(parkingLot)

	if len(attendant.assignedLots) == 1 && attendant.assignedLots[0] == parkingLot {
		t.Logf("Assigned lot added successfully")
	} else {
		t.Error("Failed to add assigned lot")
	}
}
