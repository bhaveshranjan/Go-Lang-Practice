package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	fmt.Println("Current time: ", currentTime)
	fmt.Printf("Type of currentTime %T\n", currentTime)

	//output = Current time:  2024-06-05 08:17:49.550967 +0530 IST m=+0.000262697
	// Type of currentTime time.Time

	// formatted := currentTime.Format("dd:mm:yyyy, hh:mm:ss")
	// Go lang don't follow this format they will print you this as string

	formatted := currentTime.Format("02-01-2006, 15:04:05")
	fmt.Println("Formatted time: ", formatted)

	//Formatted time:  05-06-2024, 08:17:49

	formatted2 := currentTime.Format("2006/01/02, 3:04 PM")
	fmt.Println("Formatted time 2: ", formatted2)

	//Formatted time 2:  2024/06/05, 8:22 AM

	formatted3 := currentTime.Format("2006/01/02, Monday")
	fmt.Println("Formatted time 3: ", formatted3)

	//Formatted time 3:  2024/06/05, Wednesday

	layout_str := "2006-01-02"
	dateStr := "2023-11-25"
	formatted_time, _ := time.Parse(layout_str, dateStr)
	fmt.Println("Formatted time: ", formatted_time)

	//Formatted time:  2023-11-25 00:00:00 +0000 UTC

	//add 1 more day in currentTime

	new_date := currentTime.Add(48 * time.Hour)
	fmt.Println("new_date time: ", new_date)
	formatted_new_date := new_date.Format("2006/01/02 Monday")
	fmt.Println("formatted_new_date time: ", formatted_new_date)

	//formatted_new_date time:  2024/06/07 Friday
}
