package main

import "fmt"

func main() {
	normalBuilder := getBuilder()

	director := newDirector(normalBuilder)
	sportCar := director.buildCar()

	fmt.Println(sportCar)
}
