package main

type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) buildCar() Car {
	d.builder.setDoor()
	d.builder.setWindows()
	return d.builder.getCar()
}
