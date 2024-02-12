package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ParkingLot struct {
	lotID        int
	capacity     int
	parkingSlots map[int]Car
}

func NewParkingLot(capacity int) (*ParkingLot, error) {
	if capacity < 0 {
		return nil, fmt.Errorf("Invalid capacity. Capacity must be a positive value.")
	}
	return &ParkingLot{
		lotID:        generateRandomLotID(),
		capacity:     capacity,
		parkingSlots: make(map[int]Car),
	}, nil
}

func generateRandomLotID() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(1000)
}

func (lot *ParkingLot) ParkCar(car Car) error {
	if lot.isCarParked(car) {
		return fmt.Errorf("Car with registration number is already parked")
	}

	if lot.isFull() {
		return fmt.Errorf("Parking lot is full. Trying the next one.")
	}

	nextAvailableSlot := getNextAvailableSlot(lot)
	lot.parkingSlots[nextAvailableSlot] = car

	fmt.Printf("Car parked successfully at slot %d in lot %d\n", nextAvailableSlot, lot.lotID)
	return nil
}

func (lot *ParkingLot) UnparkCar(ticket ParkingTicket) (Car, error) {
	for slot, parkedCar := range lot.parkingSlots {
		if parkedCar.parkingTicket.registrationNumber == ticket.registrationNumber {
			delete(lot.parkingSlots, slot)
			fmt.Printf("Car with registration number %s unparked.\n", ticket.registrationNumber)
			return parkedCar, nil
		}
	}

	return Car{}, fmt.Errorf("Car with registration number %s not found", ticket.registrationNumber)
}

func (lot *ParkingLot) isCarParked(car Car) bool {
	for _, parkedCar := range lot.parkingSlots {
		if parkedCar.parkingTicket.registrationNumber == car.parkingTicket.registrationNumber {
			return true
		}
	}
	return false
}

func (lot *ParkingLot) isFull() bool {
	return len(lot.parkingSlots) >= lot.capacity
}

func getNextAvailableSlot(lot *ParkingLot) int {
	rand.Seed(time.Now().UnixNano())
	for {
		slot := rand.Intn(lot.capacity)
		if _, exists := lot.parkingSlots[slot]; !exists {
			return slot
		}
	}
}

func (lot *ParkingLot) getLotID() int {
	return lot.lotID
}
