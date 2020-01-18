package main

type VehicleFactory interface {
	NewVehicle(int) (Vehiche, error)
}
