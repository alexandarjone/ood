package main

import (
	"errors"
	"fmt"
)

type SpotType int
const (
    Small SpotType = iota
    Medium
    Large
)

type ParkingSpotType interface {
    // GetType returns the type of the ParkingSpot
    GetType() SpotType
    // GetFee returns the total fee from start to end time
    GetFee(int, int) int
}

type SmallSpot struct {
    rate int
}

func (s SmallSpot) GetType() SpotType {
    return Small
}

func (s SmallSpot) GetFee(start, end int) int {
    return (end-start+1) * s. rate
}

func NewSmallSpot(rate int) ParkingSpotType {
    return &SmallSpot{rate: rate}
}

type MediumSpot struct {
    rate int
}

func (s MediumSpot) GetType() SpotType {
    return Medium
}

func (s MediumSpot) GetFee(start, end int) int {
    return (end-start+1) * s. rate
}

func NewMediumSpot(rate int) ParkingSpotType {
    return &MediumSpot{rate: rate}
}

type LargeSpot struct {
    rate int
}

func (s LargeSpot) GetType() SpotType {
    return Large
}

func (s LargeSpot) GetFee(start, end int) int {
    return (end-start+1) * s. rate
}

func NewLargeSpot(rate int) ParkingSpotType {
    return &LargeSpot{rate: rate}
}

type SpotStatus int
const (
    Available SpotStatus = iota
    Occupied
    Unavailable
)

type ParkingSpot interface {
    GetType() ParkingSpotType
    GetStatus() SpotStatus
    SetStatus(SpotStatus)
}

type VehicleParkingSpot struct {
    spotType ParkingSpotType
    status SpotStatus
}

func (s VehicleParkingSpot) GetType() ParkingSpotType {
    return s.spotType
}

func (s VehicleParkingSpot) GetStatus() SpotStatus {
    return s.status
}

func (s *VehicleParkingSpot) SetStatus(newStatus SpotStatus) {
    s.status = newStatus
}

func NewVehicleParkingSpot(spotType ParkingSpotType) ParkingSpot {
    return &VehicleParkingSpot{
        spotType: spotType,
        status: Available,
    }
}

type ParkingLot interface {
    GetSpots() []ParkingSpot
}

type SingleFloorParkingLot struct {
    spots []ParkingSpot
}

func (s SingleFloorParkingLot) GetSpots() []ParkingSpot {
    return s.spots
}

func NewSingleFloorParkingLot(typeCount []struct{spotType ParkingSpotType; count int}) ParkingLot {
    spots := []ParkingSpot{}
    for _, config := range typeCount {
        for range config.count {
            spots = append(spots, NewVehicleParkingSpot(config.spotType))
        }
    }
    return &SingleFloorParkingLot{spots: spots}
}

type Ticket interface {
    GetStartTime() int
    GetSpot() ParkingSpot
    GetPrice(int) int
}

type ParkingTicket struct {
    startTime int
    spot ParkingSpot
}

func (p ParkingTicket) GetStartTime() int {
    return p.startTime
}

func (p ParkingTicket) GetSpot() ParkingSpot {
    return p.spot
}

func (p ParkingTicket) GetPrice(currentTime int) int {
    return p.spot.GetType().GetFee(p.startTime, currentTime)
}

func NewParkingTicket(currentTime int, spot ParkingSpot) Ticket {
    ticket := &ParkingTicket{
        startTime: currentTime,
        spot: spot,
    }
    spot.SetStatus(Occupied)
    return ticket
}

type ParkingLotManager interface {
    // Park parks the car into a specific ParkingSpotType
    Park(ParkingSpotType, int) (Ticket, error)
    // Leave checks out the car and returns the total fee
    Leave(Ticket, int) int
}

type VehicleParkingLotManager struct {
    parkingLot ParkingLot
}

func (v VehicleParkingLotManager) Park(spotType ParkingSpotType, currentTime int) (Ticket, error) {
    for _, spot := range v.parkingLot.GetSpots() {
        if spot.GetType() == spotType && spot.GetStatus() == Available {
            ticket := NewParkingTicket(currentTime, spot)
            return ticket, nil
        }
    }
    return nil, errors.New("no parking spot for the requested type")
}

func (v VehicleParkingLotManager) Leave(ticket Ticket, currentTime int) int {
    price := ticket.GetPrice(currentTime)
    ticket.GetSpot().SetStatus(Available)
    return price
}

func NewVehicleParkingLotManager(parkingLot ParkingLot) ParkingLotManager {
    return &VehicleParkingLotManager{parkingLot: parkingLot}
} 

func main() {
    smallSpot := NewSmallSpot(5)
    mediumSpot := NewMediumSpot(10)
    largeSpot := NewLargeSpot(15)
    config := []struct{spotType ParkingSpotType; count int}{
        {smallSpot, 5},
        {mediumSpot, 5},
        {largeSpot, 0},
    }
    parkingLot := NewSingleFloorParkingLot(config)
    parkingLotManager := NewVehicleParkingLotManager(parkingLot)
    ticket, err := parkingLotManager.Park(smallSpot, 1)
    fmt.Println(ticket, err)
    price := parkingLotManager.Leave(ticket, 2)
    fmt.Println(price)
    ticket, err = parkingLotManager.Park(largeSpot, 1)
    fmt.Println(ticket, err)
}