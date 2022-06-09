package main

type mercedesCarFactory struct {
}

func (bmw *mercedesCarFactory) createSportCar() iCar {
	return &car{
		logo:       "mercedes",
		carType:    "sport",
		engineSize: 5,
		seats:      2,
	}
}

func (bmw *mercedesCarFactory) createSUVCar() iCar {
	return &car{
		logo:       "mercedes",
		carType:    "SUV",
		engineSize: 5,
		seats:      7,
	}
}
