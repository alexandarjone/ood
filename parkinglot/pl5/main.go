package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// --- Parking Lot System Interface ---
type SpotType int
const (
	Small SpotType = iota
	Medium
	Large
)

type PriceModel interface {
	GetPrice(startTime, endTime time.Time) float64
}

type SpotStatus int
const (
	Available SpotStatus = iota
	Occupied
)
type Spot interface {
	GetType() SpotType
	GetPriceModel() PriceModel
	GetStatus() SpotStatus
	Park() error
	Leave() error
}

type TicketStatus int
const (
	Unpaid TicketStatus = iota
	Paid	
)

type Ticket interface {
	GetStartTime() time.Time
	GetSpot() Spot
	Checkout(endTime time.Time) error
	GetStatus() TicketStatus
}

type ParkingLot interface {
	GetSpots(SpotType) []Spot
	ParkIntoSpot(Spot) (Ticket, error)
	Checkout(Ticket) error
}

// --- Parking Lot System Implementation
// Price Model
type flatRatePriceModel struct {
	// per 15-minute rate
	rate float64
}
func (f *flatRatePriceModel) GetPrice(startTime, endTime time.Time) float64 {
	duration := (endTime.Sub(startTime) + 15 * time.Minute) / 15 * time.Minute
	return float64(duration) * f.rate
}
func NewFlatRatePriceModel(rate float64) PriceModel {
	return &flatRatePriceModel{rate: rate}
}
// Spot
type spot struct {
	spotType SpotType
	priceModel PriceModel
	status SpotStatus
	lock sync.Mutex
}
func (s *spot) GetType() SpotType {
	return s.spotType
}
func (s *spot) GetPriceModel() PriceModel {
	return s.priceModel
}
func (s *spot) GetStatus() SpotStatus {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.status
}
func (s *spot) Park() error {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.status != Available {
		return errors.New("the spot is not available right now")
	}
	s.status = Occupied
	return nil
}
func (s *spot) Leave() error {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.status != Occupied {
		return errors.New("the spot is not occupied right now")
	}
	s.status = Available
	return nil
}
func NewSpot(spotType SpotType, priceModel PriceModel) Spot {
	return &spot{
		spotType: spotType,
		priceModel: priceModel,
		status: Available,
	}
}
// Ticket
type ticket struct {
	startTime time.Time
	spot Spot
	status TicketStatus
	lock sync.Mutex
}
func (t *ticket) GetStartTime() time.Time {
	return t.startTime
}
func (t *ticket) GetSpot() Spot {
	return t.spot
}
func (t *ticket) Checkout(endTime time.Time) error {
	t.lock.Lock()
	defer t.lock.Unlock()
	// call 3rd party api
	fmt.Println("paid:", t.spot.GetPriceModel().GetPrice(t.startTime, endTime))
	t.status = Paid
	return nil
}
func (t *ticket) GetStatus() TicketStatus {
	t.lock.Lock()
	defer t.lock.Unlock()
	return t.status
}
func NewTicket(startTime time.Time, spot Spot) Ticket {
	return &ticket{
		startTime: startTime,
		spot: spot,
		status: Unpaid,
	}
}
// ParkingLot
type parkingLot struct {
	parkingSpots map[SpotType][]Spot
}
func (p *parkingLot) GetSpots(spotType SpotType) []Spot {
	return p.parkingSpots[spotType]
}
func (p *parkingLot) ParkIntoSpot(spot Spot) (Ticket, error) {
	err := spot.Park()
	if err != nil {
		return nil, fmt.Errorf("failed to park into the spot: %w", err)
	}
	newTicket := NewTicket(time.Now(), spot)
	return newTicket, nil
}
func (p *parkingLot) Checkout(ticket Ticket) error {
	spot := ticket.GetSpot()
	err := spot.Leave()
	if err != nil {
		return fmt.Errorf("failed to checkout: %w", err)
	}
	err = ticket.Checkout(time.Now())
	if err != nil {
		return fmt.Errorf("failed to checkout: %w", err)
	}
	return nil
}
func CreateNewFlatRateSpots(spotType SpotType, count int, rate float64) []Spot {
	newSpots := []Spot{}
	priceModel := NewFlatRatePriceModel(rate)
	for range count {
		newSpots = append(newSpots, NewSpot(spotType, priceModel))
	}
	return newSpots
}
func NewParkingLot(
	smallCount, mediumCount, largeCount int,
	smallRate, mediumRate, largeRate float64,
) ParkingLot {
	spotMap := make(map[SpotType][]Spot)
	spotMap[Small] = CreateNewFlatRateSpots(Small, smallCount, smallRate)
	spotMap[Medium] = CreateNewFlatRateSpots(Medium, mediumCount, mediumRate)
	spotMap[Large] = CreateNewFlatRateSpots(Large, largeCount, largeRate)
	return &parkingLot{parkingSpots: spotMap}
}