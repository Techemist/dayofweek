package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Days int

const (
	Mon Days = iota
	Tue
	Wed
	Thu
	Fri
	Sat
	Sun
)

func main() {
	fmt.Print("Please Enter the day of the week (Mon,Tue,Wed,Thu,Fri,Sat,Sun): ")
	reader := bufio.NewReader(os.Stdin)
	var str string
	str, _ = reader.ReadString('\n')
	// Since the user might enter the day in many ways below three line of code
	//standardize the input
	str = strings.ToLower(str)
	str = strings.TrimSpace(str)
	str = strings.Title(str)
	fmt.Print("Please Enter the number of days after that(0-500): ")
	var step int
	_, err := fmt.Scanf("%d", &step)
	if err != nil {
		log.Fatal(err)
	}
	if step < 0 || step > 500 {
		fmt.Print("Out of range! please try again")
		return
	}
	switch str { // Based on user's input the function nextDay would call
	case "Mon":
		fmt.Print(nextDay(Mon, step).String())
	case "Tue":
		fmt.Print(nextDay(Tue, step).String())
	case "Wed":
		fmt.Print(nextDay(Wed, step).String())
	case "Thu":
		fmt.Print(nextDay(Thu, step).String())
	case "Fri":
		fmt.Print(nextDay(Fri, step).String())
	case "Sat":
		fmt.Print(nextDay(Sat, step).String())
	case "Sun":
		fmt.Print(nextDay(Sun, step).String())
	default:
		fmt.Println("Not Valid Days Entered")
	}
}
func (d Days) String() string {
	return [...]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}[d]
}

/*
This function receives two input. First is the day and second is the step
that we want to move forward
*/
func nextDay(cd Days, s int) Days {
	// Since After each 7 days we will reach to the same day the reminder of the step in respect to 7 has the same effect
	temp := (int(cd) + s) % 7
	return Days(temp)
}
