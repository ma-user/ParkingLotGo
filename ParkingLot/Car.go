package main

import (
	"fmt"
	"regexp"
)

type Car struct {
	parkingTicket ParkingTicket
	color         string
}

type ParkingTicket struct {
	registrationNumber RegistrationNumber
}

type RegistrationNumber string

func NewCar(registrationNumber RegistrationNumber, color string) (*Car, error) {
	if !isValidRegistrationNumber(string(registrationNumber)) {
		return nil, fmt.Errorf("Invalid registration number format")
	}

	return &Car{
		parkingTicket: ParkingTicket{registrationNumber: registrationNumber},
		color:         color,
	}, nil
}

func isValidRegistrationNumber(registrationNumber string) bool {
	pattern := "[A-Z]{3}\\d{3}"
	compiledPattern := regexp.MustCompile(pattern)
	return compiledPattern.MatchString(registrationNumber)
}

func (car *Car) getColor() string {
	return car.color
}
