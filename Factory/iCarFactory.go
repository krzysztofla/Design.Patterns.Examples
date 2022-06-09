package main

import "fmt"

type iCarFactory interface {
	createSportCar() iCar
	createSUVCar() iCar
}

func getCarFactory(brand string) (iCarFactory, error) {
	if brand == "bmw" {
		return &bmwCarFactory{}, nil
	}
	if brand == "mercedes" {
		return &mercedesCarFactory{}, nil
	}
	return nil, fmt.Errorf("wrong factory type")
}
