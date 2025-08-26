package main

import (
	"fmt"
	t "time"
)

func main() {
	var fullName string
	var age int
	var sex bool
	var number string
	var color string
	// declearing my variables
	fullName = "Naomi Peter"
	age = 24 + 354
	//is she female
	sex = true
	//her phone number
	number = "07060909544"
	//her color
	color = "chocolate"
	// let print the information above
	fmt.Println("fullNane", fullName)
	fmt.Println(age)
	fmt.Println(sex)
	fmt.Println(number)
	fmt.Println(color)
	fmt.Println(t.Now())
}
