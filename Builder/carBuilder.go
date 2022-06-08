package main

type carBuilder struct {
	carWindows string
	carDoors   string
}

func (cb *carBuilder) setDoor() {
	cb.carDoors = "new tinted car windows"
}

func (cb *carBuilder) setWindows() {
	cb.carDoors = "new sport car doors"
}

func (cb *carBuilder) getCar() Car {
	return Car{
		windows: cb.carWindows,
		dor:     cb.carDoors,
	}
}
