package main

import (
	"fmt"
	t "time"
)

func main() {
	var myPhones string
	var Battery int
	var Product string
	var price float64
	var purchase bool

	// declaring my variables
	myPhones = "android"
	Battery = 75
	Product = "tecno"
	price = 34.000
	purchase = true

	// printing my phone details
	fmt.Println("Product:", Product)
	fmt.Println("Phone Type:", myPhones)
	fmt.Println("Battery Percentage:", Battery, "%")
	fmt.Println("Price: $", price)
	fmt.Println("Purchased:", purchase)
	fmt.Println("Current Time:", t.Now())
}
