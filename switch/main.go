package main

import "fmt"

func main() {
	day := 2

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Unknown Day")

	}
	month := "March"

	switch month {
	case "January", "February", "March":
		fmt.Println("Winter")
	case "April", "May", "June":
		fmt.Println("Summer")
	case "July", "August", "September":
		fmt.Println("Spring")
	default:
		fmt.Println("Unknown Season")

	}

	temp := 25

	switch {
	case temp < 0:
		fmt.Println("Freezing")
	case temp >= 0 && temp < 10:
		fmt.Println("Cold")
	case temp >= 10 && temp < 10:
		fmt.Println("Cool")
	case temp >= 20 && temp < 30:
		fmt.Println("Warm")
	default:
		fmt.Println("Unknown Day")

	}
}
