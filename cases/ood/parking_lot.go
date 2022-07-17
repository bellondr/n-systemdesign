package main

func main (){

}

type VehicleType string

type ParkingLot struct {
	UnusedParkingSpots []ParkingSpot
	UsedParkingSpots []ParkingSpot
}

type Vehicle struct {
	VehicleType string
	Id string
}

type ParkingSpot struct {
	VehicleType string
	Id string
}

type Ticket struct {
	VehicleType string
	Id string
}