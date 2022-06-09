package main

import "fmt"

func main() {
	carFactory, _ := getCarFactory("bmw")
	suv := carFactory.createSUVCar()
	fmt.Println(suv.getLogo())

}
