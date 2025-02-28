/*
	requirement:
		1. locker can have two small packages or one large package
		2. package: customer id
		3. one customer - one locker
		4. allocate package into locker
	clarification:
		1. need to handle race condition
		2. no other package type
		3. verify id when retrieving:
			i. assign a otp to locker
		4. assume package status transition perfectly:
			i. being delivered
			ii. in locker
			iii. taken
		5. only two sizes
		6. each locker is same size, same location
		7. put package:
			i. put one package
			ii. put another package

    object clear, clarify object,
    function (edge cases): only include minimum use cases
    clarify features when defining classes
    use interface when time is enough
    for simplicity, let's ignore getters
    for verifier, return nil
*/

package main

import "errors"

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

// TODO: directly call on locker
type LockerCenter interface {
	GetAllAvailableLockers() []Locker
	// no need to add this, can have LockerManager that automate this process
	PutPackage(PackageItem) (Ticket, error)
	// TakeAllItem takes locker id, otp and unlock and clear the locker
	TakeAllItem(int64, string) error
	getLockerID()
}

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

func NewTicket(packageItem PackageItem, locker Locker) Ticket {
	return &ticket{
		packageItem: packageItem,
		locker: locker,
	}
}

// Locker implementation
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

func (l *locker) GetPassCode() string {
	return l.passcode
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
	l.packageItems = append(l.packageItems, newPackage)
	newTicket := NewTicket(newPackage, l)

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
	return nil
}
