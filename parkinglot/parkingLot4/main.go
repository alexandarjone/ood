package main

import "fmt"

type ParkingSpotType int
const (
	Small ParkingSpotType = iota
	Large
)

type Location interface {
	GetLatitude() float64
	GetLongitude() float64
	GetAltitude() float64
}

type ParkingSpot interface {
	GetLocation() Location
}

type ParkingLot interface {
	GetParkingSpotByID(int64) ParkingSpot
	GetAvailableParkingSpotsByType(ParkingSpotType) []ParkingSpot
}

type parkingLot struct {
	// parking spot id => parkingSpot
	parkingSpot map[int64]ParkingSpot
}

type Terminal interface {
	GetLocation() Location
}

type AssignStrategy interface {
	GetBestSpot(Terminal, ParkingSpotType)
}

type minParkingSpotHeap []ParkingSpot
type locationAssignStrategy struct {
	// parkingSpotsHeaps map[ParkingSpotType]map[Terminal]minParkingSpotHeap
	// parkingSpotsSet map[ParkingSpotType]map[Terminal]map[minParkingSpotHeap]bool
	parkingLot ParkingLot
}


func (s *locationAssignStrategy) GetBestSpot(terminal Terminal, spotType ParkingSpotType) {
	for _, spot := range s.parkingLot.GetAvailableParkingSpotsByType(spotType) {
		fmt.Println(spot)
	}
}

func NewLocationAssignStrategy(parkingLot ParkingLot) AssignStrategy {
	return &locationAssignStrategy{
		parkingLot: parkingLot,
	}
}

/*
when exit, push spot back to all heaps:
	1. spot is already in the heap:

	2. spot is not in the heap:
		push it back




3 terminal

3 spotType

9 heap
(terminalA, Small)





A








B (spotB, spotD, spotZ, spotA)


	parking_spot


	-----------------------
	|
	|
	|
	|
	|
	|

*/