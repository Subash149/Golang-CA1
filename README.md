# Flight Reservation Directory

A small Go package that models a simple in-memory flight reservation system —
scheduling flights, looking them up, reserving seats, and cancelling routes.

Built as a learning project to practice Go fundamentals: structs, pointer
receivers, maps, and idiomatic error handling.

## Domain

**Aviation** — a `Flight` represents one scheduled route between two airports.

## Project Structure

```
Golang-CA1/
├── go.mod
├── main.go              # demo program using the airline package
└── airline/
    └── airline.go        # Flight struct + Directory (core logic)
```

## The `Flight` struct

```go
type Flight struct {
    FlightNumber string
    Origin       string
    Destination  string
    Capacity     int
    BookedSeats  int
}
```

## API

All operations live on a `*Directory`, created with `airline.NewDirectory()`.

| Method | Description |
|---|---|
| `ScheduleFlight(f Flight) error` | Adds a new flight. Fails if origin == destination, capacity <= 0, flight number is empty, or the flight number is already scheduled. |
| `FindFlight(flightNum string) (*Flight, error)` | Looks up a flight by flight number. Returns an error if not found. |
| `ReserveSeat(flightNum string, count int) error` | Books `count` seats on a flight. Fails if the count is non-positive or would exceed capacity. |
| `CancelFlightRoute(flightNum string) error` | Removes a scheduled flight. Fails if the flight number doesn't exist. |

## Getting Started

### Prerequisites

- Go 1.21 or later

### Run the demo

```bash
git clone https://github.com/tamilselvan-v07/Golang-CA1.git
cd Golang-CA1
go run main.go
```

Expected-ish output:

```
Scheduled AI202 successfully
Expected error scheduling AI999: origin and destination cannot be the same
Found flight: {FlightNumber:AI202 Origin:Chennai Destination:Delhi Capacity:3 BookedSeats:0}
Reserved 2 seats on AI202
Expected error reserving 5 seats: cannot reserve 5 seat(s) on flight AI202: only 1 available
AI202 after reservation: {FlightNumber:AI202 Origin:Chennai Destination:Delhi Capacity:3 BookedSeats:2}
Cancelled AI202 successfully
Expected error after cancel: flight AI202 not found
```

### Use it in your own code

```go
import "flightres/airline"

dir := airline.NewDirectory()

err := dir.ScheduleFlight(airline.Flight{
    FlightNumber: "AI202",
    Origin:       "Chennai",
    Destination:  "Delhi",
    Capacity:     180,
})

err = dir.ReserveSeat("AI202", 4)

flight, err := dir.FindFlight("AI202")

err = dir.CancelFlightRoute("AI202")
```

## Design Notes

- Flights are stored as `*Flight` (pointers) inside a `map[string]*Flight`,
  so that `ReserveSeat` can mutate the stored flight's `BookedSeats` in
  place rather than working on a disconnected copy.
- Errors are returned, not panicked — callers decide how to handle a
  failed reservation or a missing flight.
