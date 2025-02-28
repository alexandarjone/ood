/*
    airline management requirement
        1. customer search for light: day, from and to
        2. reserve ticket
        3. check flight schedule, departure time, available seats, arriving time
        4. reserve for other passengers for one flight
    
    clarification
        1. planes, flights are preconfigues
        2. customer info, payment is handled
        3. don't need to handle race condition
    
    workflow:
        1. user input: date, source, destination and get a list of flight
        2. flight: departure time, available seats, arriving time
        3. user reserve a flight for himself / other people
*/

type User interface {
    GetID() int64
}

type Plane interface {
    GetID() int64
}

type SeatStatus
const (
    Available SeatStatus = iota
    Reserved
    Occupied
    Unvavailable
)
type Seat interface {
    GetID() int64
    GetStatus() SeatStatus
}

type Flight interface {
    GetDepartureTime() time.Time
    GetSource() string
    GetDetination() string
    GetAvailableSeats() []Seat
    ReserveSeat(int64)
}

type FlightCenter interface {
    GetFlight(int64) Flight
    GetFlights(FlightFilter) []Flight
}

type FlightFilter interface {
    IsValid(Flight) bool
}

type FlightOrder interface {
    GetCustomerID() int64
    GetPassengerID() int64
    GetFlightID() int64
}

type OrderCenter interface {
    // ReserveFlight takes user id, flight id, seat id and book the flight
    ReserveFlight(int64, int64, int64) error
    // ReserveFlight takes customer id and returns all their flight orders
    GetFlightOrders(int64) []FlightOrder
}

// FlightCenter implemention
type flightCenter struct {
    // flight id => Flight
    flights map[int64]Flight
}

func (f *flightCenter) GetFlights(filter FlightFilter) []Flight {
    filteredFlights := []Flight{}
    for _, flight := range f.flights {
        if filter.IsValid(flight) {
            filteredFlights = append(filteredFlights, flight)
        }
    }
    return result
}
// TODO: constructor

// OrderCenter implemention
type orderCenter struct {
    // order id => Order
    orders map[int64]Order
    // dependency injection
    flightCenter FlightCenter
}

// ReserveFlight takes user id, flight id, seat id and book the flight
func (o *orderCenter) ReserveFlight(userID int64, passengerID int64, flightID int64, seatID int64) error {
    targetFlight := o.flghtCenter.GetFlight(flightID)
    availableSeats := targetFlight.GetAvailableSeats()
    isAvailable := false
    for _, availableSeat := range availableSeats {
        if availableSeat.GetID() == seatID {
            isAvailable = true
        }
    }
    // TODO: handle payment
    targetFlight.ReserveSeat(seatID)
    newOrder := NewFlightOrder(userID, passengerID, flightID, seatID)
    o.orders[newOrder.GetID()] = newOrder
}
// ReserveFlight takes customer id and returns all their flight orders
GetFlightOrders(int64) []FlightOrder

func main() {
    fmt.Printf("Hello LeetCoder")
}