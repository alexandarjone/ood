package main

import "sync"

type ParkingSpotType interface {
	GetName() string
	GetRate() int
}

type ParkingSpot interface {
	GetType() ParkingSpotType
	GetStatus() string
	SetStatus(string)
	GetLocation() string
}

type Ticket interface {
	GetStartTime() int
	GetParkingSpot() ParkingSpot
}

type ParkingLot interface {
	IssueTicket(ParkingSpotType, int) Ticket
	AcceptTicket(Ticket) bool
}

type GarageParkingSpotType struct {
	name string
	rate int
}

func (g GarageParkingSpotType) GetName() string {
	return g.name
}

func (g GarageParkingSpotType) GetRate() int {
	return g.rate
}

func NewGarageParkingSpotType(name string, rate int) ParkingSpotType {
	return &GarageParkingSpotType{
		name: name,
		rate: rate,
	}
}

type GarageParkingSpot struct {
	spotType ParkingSpotType
	status string
	location string
	lock sync.Mutex
}

func (g *GarageParkingSpot) GetType() ParkingSpotType {
	return g.spotType
}

func (g *GarageParkingSpot) GetStatus() string {
	g.lock.Lock()
	defer g.lock.Unlock()
	status := g.status
	return status
}

func (g *GarageParkingSpot) SetStatus(status string) {
	g.lock.Lock()
	defer g.lock.Unlock()
	g.status = status
}

func (g *GarageParkingSpot) GetLocation() string {
	return g.location
}

func NewGarageParkingSpot(spotType ParkingSpotType, location string) ParkingSpot {
	return &GarageParkingSpot{
		spotType: spotType,
		status: "available",
		location: location,
	}
}

type GarageTicket struct {
	spot ParkingSpot
	startTime int
}

func (g GarageTicket) GetStartTime() int {
	return g.startTime
}

func (g GarageTicket) GetParkingSpot() ParkingSpot {
	return g.spot
}

func NewGarageTicket(spot ParkingSpot, startTime int) Ticket {
	return &GarageTicket{
		spot: spot,
		startTime: startTime,
	}
}

type GarageParkingLot struct {
	parkingSpots []ParkingSpot
	paymentProcessor Payment
}

func (g GarageParkingLot) IssueTicket(spotType ParkingSpotType, time int) Ticket {
	for _, spot := range g.parkingSpots {
		if spot.GetStatus() == "available" && spot.GetType() == spotType {
			spot.SetStatus("taken")
			return NewGarageTicket(spot, time)
		}
	}
	return nil
}

func (g GarageParkingLot) AcceptTicket(ticket Ticket) bool {
	ticket.GetParkingSpot().SetStatus("available")
	return g.paymentProcessor.Process(ticket)
}

type Payment interface {
	Process(Ticket) bool
}

type ParkingLotSystem interface {
	GetParkingLots() []ParkingLot
	AddParkingLot(ParkingLot)
}

type GarageParkingLotSystem struct {
	parkingLots []ParkingLot
}

func (g GarageParkingLotSystem) GetParkingLots() []ParkingLot {
	return g.parkingLots
}

func (g *GarageParkingLotSystem) AddParkingLot(parkingLot ParkingLot) {
	g.parkingLots = append(g.parkingLots, parkingLot)
}

func NewGarageParkingLotSystem(parkingLots []ParkingLot) ParkingLotSystem {
	return &GarageParkingLotSystem{
		parkingLots: parkingLots,
	}
}