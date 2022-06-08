package main

type iBuilder interface {
	setDoor()
	setWindows()
	getCar() Car
}

func getBuilder() iBuilder {
	return &carBuilder{}
}
