package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type SpotType int
const (
	SmallSpot SpotType = iota
	MediumSpot
	LargeSpot
	LargeElectricChargerSpot
)

type PriceModel interface {
	GetFee(time.Time, time.Time) float64
}

type flatRatePriceModel struct {
	rate float64
}

func (f *flatRatePriceModel) GetFee(start, end time.Time) float64 {
	duration := ((end.Sub(start) + 15 * time.Minute) / (15 * time.Minute))
	return float64(duration) * f.rate
}

func NewFlatRatePriceModel(rate float64) PriceModel {
	return &flatRatePriceModel{
		rate: rate,
	}
}

type Status int
const (
	Available Status = iota
	Occupied
	Unavailable
)

type ParkingSpot interface {
	GetParkingSpotType() SpotType
	GetPriceModel() PriceModel
	GetStatus() Status
	Park() error
	Leave() error
}

type parkingSpot struct {
	spotType SpotType
	priceModel PriceModel
	status Status
	lock sync.Mutex
}

func (p *parkingSpot) GetParkingSpotType() SpotType {
	return p.spotType
}

func (p *parkingSpot) GetPriceModel() PriceModel {
	return p.priceModel
}

func (p *parkingSpot) GetStatus() Status {
	return p.status
}

func (p *parkingSpot) Park() error {
	p.lock.Lock()
	defer p.lock.Unlock()
	if p.status != Available {
		return errors.New("the spot is unavailable")
	}
	p.status = Occupied
	return nil
}

func (p *parkingSpot) Leave() error {
	p.lock.Lock()
	defer p.lock.Unlock()
	if p.status != Occupied {
		return errors.New("the spot is not occupied")
	}
	p.status = Available
	return nil
}

func NewParkingSpot(spotType SpotType, priceModel PriceModel) ParkingSpot {
	return &parkingSpot{
		spotType: spotType,
		priceModel: priceModel,
	}
}

type Ticket interface {
	GetParkingSpot() ParkingSpot
	GetStartTime() time.Time
	GetFee(time.Time) float64
}

type ticket struct {
	parkingSpot ParkingSpot
	startTime time.Time
}

func (t *ticket) GetParkingSpot() ParkingSpot {
	return t.parkingSpot
}

func (t *ticket) GetStartTime() time.Time {
	return t.startTime
}

func (t *ticket) GetFee(endTime time.Time) float64 {
	priceModel := t.parkingSpot.GetPriceModel()
	return priceModel.GetFee(t.startTime, endTime)
}

func NewTicket(parkingSpot ParkingSpot, startTime time.Time) Ticket {
	return &ticket{
		parkingSpot: parkingSpot,
		startTime: startTime,
	}
}

type ParkingLot interface {
	GetParkingSpots(SpotType) []ParkingSpot
	Park(ParkingSpot) (Ticket, error)
	Checkout(Ticket) (float64, error)
}

type parkingLot struct {
	parkingSpots map[SpotType][]ParkingSpot
}

func (p *parkingLot) GetParkingSpots(spotType SpotType) []ParkingSpot {
	availableSpots := []ParkingSpot{}
	for _, spot := range p.parkingSpots[spotType] {
		if spot.GetStatus() == Available {
			availableSpots = append(availableSpots, spot)
		}
	}
	return availableSpots
}

func (p *parkingLot) Park(spot ParkingSpot) (Ticket, error) {
	err := spot.Park()
	if err != nil {
		return nil, fmt.Errorf("failed to park into the spot: %w", err)
	}
	ticket := NewTicket(spot, time.Now())
	return ticket, nil
}

func (p *parkingLot) Checkout(ticket Ticket) (float64, error) {
	spot := ticket.GetParkingSpot()
	err := spot.Leave()
	if err != nil {
		return -1.0, fmt.Errorf("failed to checkout: %w", err)
	}
	fee := ticket.GetFee(time.Now())
	return fee, nil
}

func NewParkingLot(
	smallCount, mediumCount, largeCount int,
	smallRate, mediumRate, largeRate float64,
) ParkingLot {
	smallPriceModel := NewFlatRatePriceModel(smallRate)
	mediumPriceModel := NewFlatRatePriceModel(mediumRate)
	largePriceModel := NewFlatRatePriceModel(largeRate)
	spotsMap := make(map[SpotType][]ParkingSpot)
	smallParkingSpots := []ParkingSpot{}
	for range smallCount {
		smallParkingSpots = append(smallParkingSpots, NewParkingSpot(SmallSpot, smallPriceModel))
	}
	spotsMap[SmallSpot] = smallParkingSpots
	mediumParkingSpots := []ParkingSpot{}
	for range mediumCount {
		mediumParkingSpots = append(mediumParkingSpots, NewParkingSpot(MediumSpot, mediumPriceModel))
	}
	spotsMap[MediumSpot] = mediumParkingSpots
	largeParkingSpots := []ParkingSpot{}
	for range largeCount {
		largeParkingSpots = append(largeParkingSpots, NewParkingSpot(LargeSpot, largePriceModel))
	}
	spotsMap[LargeSpot] = largeParkingSpots
	return &parkingLot{parkingSpots: spotsMap}
}

func main() {
	parkingLot := NewParkingLot(10, 10, 1, 5.0, 10.0, 15.0)
	largeSpots := parkingLot.GetParkingSpots(LargeSpot)
	if len(largeSpots) > 0 {
		spot := largeSpots[0]
		largeSpots = largeSpots[1:]
		ticket, _ := parkingLot.Park(spot)
		time.Sleep(2*time.Second)
		price, _ := parkingLot.Checkout(ticket)
		fmt.Println(price)
	}
}