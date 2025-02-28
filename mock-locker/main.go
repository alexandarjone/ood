package main
/*
    amazon locker
        1. assign packages to locker at single location
        2. each package has id, customer id, size(small, large)
        3. each locker has id, can hold:
            i. 1 large
            ii. up to 2 small packages belong to the same customer
        4. given a list of packages, available lockers, function should assign packages to lockers
    
    clarification
        1. assign 2 small to same locker as much as possible
        2. no need for handling race condition
        3. user can pickup item inside locker using code (get code after purchace item)
        4. after 48 hours, code would be invalid, item goes back to package center
    
    workflow
        1. courier get a list of lockers from locker center based on a list of packages from package center
        2. courier put package(s) into locker and 
            i. locker starts a process that sleeps for 48 hours, then mark the 
                package as expire and notify the locker center's manager
            ii. get the ticket that keeps track of the relationship between 
                package and locker, and the passcode to unlock the locker.
        3. (ticket would be sent to customer) customer get the locker ID from ticket, and open
            the locker with the passcode on ticket.

    core struct:
        1. package
        2. locker
        3. package center
        4. locker center
        5. ticket

    feedback

    wayne:
        ***workflow after core struct***
        const capitalize (TODO)
    
    mengting:
    pros
        confident
        clarify
    cons
        no communication when implementing
        does not express OOP (reuse, paren-child)
*/

type PackageStatus int
const (
    BeingDelivered PackageStatus = iota
    InLocker
    Expired
    SentBack
)

type PackageSize int
const (
    Small PackageSize = iota
    Large
)

type PackageItem interface {
    GetID() int64
    GetStatus() PackageStatus
    SetStatus(PackageStatus)
    GetSize() PackageSize
    GetCutomerID() int64
}

type Password string

type LockerStatus int
const (
    Empty LockerStatus = iota
    OccupiedByOneSmall
    Full
)

type Locker interface {
    GetID() int64
    PutPackageIn(Package) (Ticket, error)
    ClearLocker(Password) error
    GetLockerStatus() LockerStatus
    GetLockedTime() time.Time
}

type HeightLocker interface {
    GetHeight()
    Locker
}

type LockerModel interface {
    Canfit(Package)
}

type Ticket interface {
    GetPackageID() int64
    GetLockerID() int64
    GetPassword() Password
}

type PackageCenter interface {
    GetAllPackages() []PackageItem
}

type LockerCenter interface {
    // assume lockers are always enough
    AssignPackages([]Package)
    StartLockerExpirationCRON()
    GetLockerByID(int64) Locker
}

// LockerCenter implementation
/*
        Empty: {
            "001": locker1,
            "002": locker2
        }
        OccupiedByOneSmall: {
            "003": locker3,
        }
        Full: {
            "004": locker4,
        }
*/
type lockerCenter struct {
    lockers map[LockerStatus]map[int64]Locker
    tickets []Ticket
}

func (l *lockerCenter) AssignPackages(newPackages []Package) {
    customerToPackages := make(map[int64][]Package)
    for _, newPackage := range newPackages {
        customerID := newPackage.GetCustomerID()
        customerToPackages[customerID] = append(customerToPackages[customerID], newPackage)
    }
    for _, pacakgeItems := range customerToPackages {
        sort.Slice(pacakgeItems, func(i, j int) bool {
            return pacakgeItems[i].GetType() == Small
        })
        index := 0
        n := len(pacakgeItems)
        for index < n {
            if pacakgeItems[index].GetType() == Small &&
                index < n-1 && pacakgeItems[index+1].GetType() == Small {
                    for 
                }
        }
    }
}

func (l *) StartLockerExpirationCRON() {
    
}

func cron() {
    // TODO: update package status to expire
    // TODO: clear locker
    // TODO: call api to notify user
}

func main() {
    fmt.Printf("Hello LeetCoder")
}
