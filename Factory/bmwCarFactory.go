package main

type bmwCarFactory struct {
}

func (bmw *bmwCarFactory) createSportCar() iCar {
	return &car{
		logo:       "bmw",
		carType:    "sport",
		engineSize: 4,
		seats:      2,
	}
}

func (bmw *bmwCarFactory) createSUVCar() iCar {
	return &car{
		logo:       "bmw",
		carType:    "SUV",
		engineSize: 3,
		seats:      5,
	}
}
