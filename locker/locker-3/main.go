package main

import (
	"errors"
	"fmt"
)

/*
workflow
	1. courier get locker id from system based on package info
	2. courier store package in locker
	3. customer gets a ticket that keeps track of:
		a. locker id
		b. package id
		c. password
	4. customer open locker with passcode

object
	1. package
	2. package manager (TODO!!!)
	3. locker
	4. locker manager
	5. ticket
	6. ticket manager
*/

type PackageSize int
const (
	Small PackageSize = 1
	Large PackageSize = 2
)

type PackageItem interface {
	GetID() int64
	GetSize() PackageSize
	GetCustomerID() int64
}

type PackageManager interface {
	GetPackageByID(int64) PackageItem
}

type Locker interface {
	GetID() int64
	GetSlotCount() int
}

type LockerManager interface {
	// AssignPackageToLocker takes locker id, package id and assign the package into the locker, then return a ticket
	AssignPackage(int64) (Ticket, error)
	// TakePackage takes locker id, passcode and clear the locker
	UnlockLocker(int64, string) error
}

type Ticket interface {
	GetTicketID() int64
	GetLockerID() int64
	GetPackageID() int64
	GetPasscode() string
}

type TicketManager interface {
	GetTicketsByLockerID(int64) []Ticket
	NewTicket(lockerID, packageID int64, passcode string) Ticket
	DeleteTicket(int64)
}

type PasswordGenerator interface {
	GeneratePassword() string
}

// LockerManager implementation
type lockerManager struct {
	// locker id => Locker
	lockers map[int64]Locker
	// customer id => []Locker
	customerIDToLockerID map[int64][]int64
	// locker it => passcode
	passcode map[int64]string
	// empty lockers
	emptyLockers map[int64]Locker
	// dependency injection
	packageManager PackageManager
	ticketManager TicketManager
	passwordGenerator PasswordGenerator
}

// canAssign takes locker id, package id and check if the package can be assigned into the locker
func (l *lockerManager) canAssign(lockerID int64, packageID int64) error {
	newPackage := l.packageManager.GetPackageByID(packageID)
	customerID := newPackage.GetCustomerID()
	slotsTaken := 0
	targetLocker, exists := l.lockers[lockerID]
	if !exists {
		return errors.New("locke does not exists")
	}
	tickets := l.ticketManager.GetTicketsByLockerID(targetLocker.GetID())
	for _, ticket := range tickets {
		packageItem := l.packageManager.GetPackageByID(ticket.GetPackageID())
		if packageItem.GetCustomerID() != customerID {
			return errors.New("the locker is occupied by another customer")
		}
		slotsTaken += int(packageItem.GetSize())
	}
	if slotsTaken + int(newPackage.GetSize()) > targetLocker.GetSlotCount() {
		return errors.New("package is too large for locker")
	}
	return nil
}

// AssignPackageToLocker takes locker id, package id and assign the package into the locker, then return a ticket
func (l *lockerManager) AssignPackageToLocker(packageID int64) (Ticket, error) {
	newPackage := l.packageManager.GetPackageByID(packageID)
	customerID := newPackage.GetCustomerID()
	for _, usedLockerID := range l.customerIDToLockerID[customerID] {
		err := l.canAssign(usedLockerID, packageID)
		if err != nil {
			continue
		}
		return l.ticketManager.NewTicket(usedLockerID, packageID, l.passcode[usedLockerID]), nil
	}
	for id := range l.emptyLockers {
		err := l.canAssign(id, packageID)
		if err != nil {
			continue
		}
		l.passcode[id] = l.passwordGenerator.GeneratePassword()
		delete(l.emptyLockers, id)
		l.customerIDToLockerID[customerID] = append(l.customerIDToLockerID[customerID], id)
		return l.ticketManager.NewTicket(id, packageID, l.passcode[id]), nil
	}
	return nil, errors.New("cannot find locker for the package")
}

func (l *lockerManager) UnlockLocker(lockerID int64, password string) error {
	storedPassword, exists := l.passcode[lockerID]
	if !exists {
		return errors.New("locker is not assigned")
	}
	if password != storedPassword {
		return errors.New("passcode is wrong")
	}
	// assume all operations are valid
	tickets := l.ticketManager.GetTicketsByLockerID(lockerID)
	packageID := tickets[0].GetPackageID()
	customerID := l.packageManager.GetPackageByID(packageID).GetCustomerID()
	// remove tickets
	for _, ticket := range tickets {
		l.ticketManager.DeleteTicket(ticket.GetTicketID())
	}
	// remove used lockers
	for index, id := range l.customerIDToLockerID[customerID] {
		if lockerID == id {
			l.customerIDToLockerID[customerID] = append(l.customerIDToLockerID[customerID][:index], 
				l.customerIDToLockerID[customerID][index+1:]...)
		}
	}
	delete(l.passcode, lockerID)
	l.emptyLockers[lockerID] = l.lockers[lockerID]
	return nil
}

func main() {
	fmt.Println(int(Large) == int(1))
}