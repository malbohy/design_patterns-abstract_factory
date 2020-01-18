package main

import (
	"fmt"
)

func main() {
	fmt.Println("abstract facotyr")

	VehicalFToryResult, _ := CreateVehicleFactory(CAR)
	microBusVehical, _ := VehicalFToryResult.NewVehicle(MICRO_BUS)
	microBusVehical.setWheels()

	fmt.Println(microBusVehical.WheelCount())

}
