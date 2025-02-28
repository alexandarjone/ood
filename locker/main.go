/*

	requirements:
	1. generate code per purchase
	2. get optimal locker for item:
		i. locker size >= item size, as close as possible
		ii. location should be as close to the user as possible

	core objects:
	1. location
	2. size
	3. locker
	4. ticket
*/

package main

import "time"

type Location interface {
	GetLatitude() float64
	GetLongitude() float64
	GetAltitude() float64
}

type Size interface {
	GetLength() float64
	GetWidth() float64
	GetHeight() float64
}

type PackageStatus int
const (
	Delivering PackageStatus = iota
	InLocker
	Picked
	Outdated
)
type PackageItem interface {
	Size
	GetID() uint64
	GetStatus() PackageStatus
	MarkDelivering() error
	MarkInLocker() error
	MarkPicked() error
	MarkOutdated() error
	GetLockerID() uint64
}

type LockerStatus int
const (
	Available LockerStatus = iota
	Occupied
	Unavailable
)
type Locker interface {
	Size
	Location
	GetID() uint64
	GetStatus() LockerStatus
	CheckIn() error
	CheckOut() error
	SetUnavailable() error
}

type Ticket interface {
	GetID() uint64
	StartTime() time.Time
	GetPackageID() uint64
	GetLockerID() uint64
}

type TicketManager interface {
	GetTicket(uint64) Ticket
	GetTicketByPackageID(uint64) Ticket
}

type LockerCenter interface {
	Location
	GetLockers() []Locker
}

type LockerCenterManager interface {
	GetOptimalLocker(PackageItem) (LockerCenter, Locker)
}