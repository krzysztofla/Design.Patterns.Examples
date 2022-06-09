package main

type iCar interface {
	setLogo(logo string)
	setType(carType string)
	setEngine(liters int)
	setSeats(seats int)
	getLogo() string
	getType() string
	getEngine() int
	getSeats() int
}

type car struct {
	logo       string
	carType    string
	engineSize int
	seats      int
}

func (c *car) setLogo(logo string) {
	c.logo = logo
}

func (c *car) setType(carType string) {
	c.carType = carType
}

func (c *car) setEngine(engineSize int) {
	c.engineSize = engineSize

}

func (c *car) setSeats(seats int) {
	c.seats = seats
}

func (c *car) getLogo() string {
	return c.logo
}

func (c *car) getType() string {
	return c.carType
}

func (c *car) getEngine() int {
	return c.engineSize
}

func (c *car) getSeats() int {
	return c.seats
}
