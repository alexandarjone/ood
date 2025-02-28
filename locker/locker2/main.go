/*
requirement:
	1. locker can have two small packages or one large package
	2. package: customer id
	3. one customer - one locker
	4. allocate package into locker

clarification:
	1. will this system get multiple requests simultaneously(race condition)?
	2. verify passcode when retrieving package
	3. pick locker randomly or under certain rules?
	4. put package:
		i. locker is empty
		ii. locker has 1 small package and the new package belongs to the same customer
*/

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type PackageType int
const (
	Small PackageType = iota
	Large
)

type PackageStatus int
const (
	BeingDelivered PackageStatus = iota
	InLocker
	PickedUp
)

type PackageItem interface {
	GetType() PackageType
	GetStatus() PackageStatus
	GetCustomerID() int64
}

type LockerStatus int
const (
	Available LockerStatus = iota
	OccupiedWithSmall
	Occupied
	Broken
)

type Locker interface {
	GetID() int64
	GetStatus() LockerStatus
	CheckPackage(PackageItem) error
	PutPackage(PackageItem) (Ticket, error)
	// TakeAllItem takes otp and unlock and clear the locker
	TakeAllItem(string) error
}

type Ticket interface {
	GetPackage() PackageItem
	GetLocker() Locker
	GetPassCode() string
}

type LockerCenter interface {
	GetAllAvailableLockers() []Locker
}

// TODO: create a LockerCenterManager to automate the process of finding a locker / clearing a locker


// PackageItem implementation
type packageItem struct {
	itemType PackageType
	status PackageStatus
	customerID int64
}

func (p *packageItem) GetType() PackageType {
	return p.itemType
}
func (p *packageItem) GetStatus() PackageStatus {
	return p.status
}
func (p *packageItem) GetCustomerID() int64 {
	return p.customerID
}

func NewPackageItem(itemType PackageType) PackageItem {
	return &packageItem{
		itemType: itemType,
		status: BeingDelivered,
	}
}

// Ticket implementation
type ticket struct {
	packageItem PackageItem
	locker Locker
	passcode string
}
func (t *ticket) GetPackage() PackageItem {
	return t.packageItem
}
func (t *ticket) GetLocker() Locker {
	return t.locker
}
func (t *ticket) GetPassCode() string {
	return t.passcode
}

func NewTicket(packageItem PackageItem, locker Locker, passcode string) Ticket {
	return &ticket{
		packageItem: packageItem,
		locker: locker,
		passcode: passcode,
	}
}

// Locker implementation
const PASSWORD_LENGTH = 8
type locker struct {
	id int64
	status LockerStatus
	passcode string
	packageItems []PackageItem
}

func (l *locker) GetID() int64 {
	return l.id
}

func (l *locker) GetStatus() LockerStatus {
	return l.status
}

func (l *locker) CheckPackage(packageItem PackageItem) error {
	if l.status == Occupied {
		return errors.New("locker is occupied")
	}
	if l.status == OccupiedWithSmall {
		if l.packageItems[0].GetCustomerID() != packageItem.GetCustomerID() {
			return errors.New("package's customer doesn't match")
		}
		if packageItem.GetType() != Small {
			return errors.New("package is too large for the locker")
		}
	}
	return nil
}

func (l *locker) PutPackage(newPackage PackageItem) (Ticket, error) {
	err := l.CheckPackage(newPackage)
	if err != nil {
		return nil, fmt.Errorf("failed to put package: %w", err)
	}

	l.packageItems = append(l.packageItems, newPackage)
	l.passcode = generatePasscode(PASSWORD_LENGTH)
	newTicket := NewTicket(newPackage, l, l.passcode)
	if l.status == Available && newPackage.GetType() == Small {
		l.status = OccupiedWithSmall
	} else {
		l.status = Occupied
	}
	return newTicket, nil
}

// TakeAllItem takes otp and unlock and clear the locker
func (l *locker) TakeAllItem(passcode string) error {
	if passcode != l.passcode {
		return errors.New("wrong passcode")
	}
	l.packageItems = []PackageItem{}
	l.status = Available
	l.passcode = ""
	return nil
}

func generatePasscode(length int) string {
	rand.Seed(time.Now().UnixNano())
	passcode := make([]byte, length)
	for i := range length {
		passcode[i] = byte('0' + rand.Intn(10))
	}
	return string(passcode)
}