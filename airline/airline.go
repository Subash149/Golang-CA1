package airline

import (
	"errors"
	"fmt"
)

type Flight struct {
	FlightNumber string
	Origin       string
	Destination  string
	Capacity     int
	BookedSeats  int
}

type Directory struct {
	flights map[string]*Flight
}

func NewDirectory() *Directory {
	return &Directory{
		flights: make(map[string]*Flight),
	}
}

func (d *Directory) ScheduleFlight(f Flight) error {
	if f.Origin == f.Destination {
		return errors.New("origin and destination cannot be the same")
	}
	if f.Capacity <= 0 {
		return errors.New("capacity must be greater than zero")
	}
	if f.FlightNumber == "" {
		return errors.New("flight number cannot be empty")
	}
	if _, exists := d.flights[f.FlightNumber]; exists {
		return fmt.Errorf("flight %s is already scheduled", f.FlightNumber)
	}

	d.flights[f.FlightNumber] = &f
	return nil
}
func (d *Directory) FindFlight(flightNum string) (*Flight, error) {
	f, ok := d.flights[flightNum]
	if !ok {
		return nil, fmt.Errorf("flight %s not found", flightNum)
	}
	return f, nil
}

func (d *Directory) ReserveSeat(flightNum string, count int) error {
	if count <= 0 {
		return errors.New("seat count to reserve must be positive")
	}

	f, err := d.FindFlight(flightNum)
	if err != nil {
		return err
	}

	if f.BookedSeats+count > f.Capacity {
		available := f.Capacity - f.BookedSeats
		return fmt.Errorf("cannot reserve %d seat(s) on flight %s: only %d available",
			count, flightNum, available)
	}

	f.BookedSeats += count
	return nil
}

func (d *Directory) CancelFlightRoute(flightNum string) error {
	if _, ok := d.flights[flightNum]; !ok {
		return fmt.Errorf("flight %s not found", flightNum)
	}
	delete(d.flights, flightNum)
	return nil
}
