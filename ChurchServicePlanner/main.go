package main

import "fmt"

func main() {
	var churchName string
	var pastorName string
	var serviceDay string
	var serviceTime string
	var numberOfAttendees int
	var isOnline bool
	//assigning my variable
	churchName = "rccg sunrise parish";
	pastorName = "pst victor omodende";
	serviceDay = "tuesday";
	serviceTime = "4:00pm";
	numberOfAttendees = 1590;
	isOnline = true;
	fmt.Println("my church", churchName, "with the speaker", pastorName, "holds every", serviceDay, serviceTime, numberOfAttendees, isOnline)

}
