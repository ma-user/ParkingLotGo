package main

import (
	"fmt"
)

type ParkingLotAttendant struct {
	assignedLots []*ParkingLot
}

func NewParkingLotAttendant() *ParkingLotAttendant {
	return &ParkingLotAttendant{
		assignedLots: make([]*ParkingLot, 0),
	}
}

func (attendant *ParkingLotAttendant) ParkCar(car Car) error {
	for _, lot := range attendant.assignedLots {
		if !lot.isFull() {
			if err := lot.ParkCar(car); err == nil {
				fmt.Printf("Car parked successfully at lot %d\n", lot.getLotID())
				return nil
			} else {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Printf("Current lot is full. Trying in next lot.")
		}
	}

	return fmt.Errorf("All parking lots are full. Cannot park the car.")
}

func (attendant *ParkingLotAttendant) assign(lot *ParkingLot) {
	attendant.assignedLots = append(attendant.assignedLots, lot)
}
