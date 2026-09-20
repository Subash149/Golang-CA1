package main

import (
	"fmt"

	"flightres/airline"
)

func readString(prompt string) string {
	var value string
	fmt.Print(prompt)
	_, err := fmt.Scanln(&value)
	if err != nil {
		return ""
	}
	return value
}

func readInt(prompt string) int {
	var value int
	fmt.Print(prompt)
	_, err := fmt.Scanln(&value)
	if err != nil {
		return 0
	}
	return value
}

func main() {
	dir := airline.NewDirectory()

	for {
		fmt.Println("\n=== Flight Reservation Directory ===")
		fmt.Println("1. Schedule Flight")
		fmt.Println("2. Find Flight")
		fmt.Println("3. Reserve Seats")
		fmt.Println("4. Cancel Flight Route")
		fmt.Println("5. Exit")

		choice := readInt("Choose an option: ")
		switch choice {
		case 1:
			flightNumber := readString("Enter flight number: ")
			origin := readString("Enter origin: ")
			destination := readString("Enter destination: ")
			capacity := readInt("Enter capacity: ")

			flight := airline.Flight{
				FlightNumber: flightNumber,
				Origin:       origin,
				Destination:  destination,
				Capacity:     capacity,
			}

			if err := dir.ScheduleFlight(flight); err != nil {
				fmt.Println("Error scheduling flight:", err)
			} else {
				fmt.Println("Flight scheduled successfully.")
			}

		case 2:
			flightNumber := readString("Enter flight number: ")
			flight, err := dir.FindFlight(flightNumber)
			if err != nil {
				fmt.Println("Error finding flight:", err)
			} else {
				fmt.Printf("Flight found: %+v\n", *flight)
			}

		case 3:
			flightNumber := readString("Enter flight number: ")
			count := readInt("Enter number of seats to reserve: ")

			if err := dir.ReserveSeat(flightNumber, count); err != nil {
				fmt.Println("Error reserving seats:", err)
			} else {
				fmt.Println("Seats reserved successfully.")
			}

		case 4:
			flightNumber := readString("Enter flight number to cancel: ")
			if err := dir.CancelFlightRoute(flightNumber); err != nil {
				fmt.Println("Error cancelling flight:", err)
			} else {
				fmt.Println("Flight route cancelled successfully.")
			}

		case 5:
			fmt.Println("Exiting flight reservation system.")
			return

		default:
			fmt.Println("Invalid option. Please choose from 1 to 5.")
		}
	}
}
