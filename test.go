package main

import (
	"bufio"     // This allows me to get input
	"fmt"       // This handles print statements
	"math/rand" // This allows me to generate random numbers
	"os"        // This helps with getting input
	"strconv"   // This allows me to convert ints to strs
	"strings"   // This helps me work with strings
	"time"      // This allows me to use time
)

func main() {
	start := time.Now() // This will record the current time (for use in measuring speed)
	fmt.Println("HELLO WORLD!")

	// Variable Declerations
	x := 12              // uint32
	y := "Wacky Tabacky" // String

	z := [5]int{2, 88, 90, 24, 5} // Int array

	// Use variables
	fmt.Println("x = " + strconv.Itoa(x))
	fmt.Println("y = " + y)
	fmt.Println("\nNow lets check out an array!")

	for i := 0; i < 5; i++ {
		fmt.Println(strconv.Itoa(i+1) + " = " + strconv.Itoa(z[i])) // This for loop shows its place and displays the arrray value
	}

	// Using functions
	detrand()

	helloWorld()

	// End: print elapsed time
	t := time.Now()
	elapsed := t.Sub(start)

	fmt.Println("\n\n==================")
	fmt.Printf("Execution Time: %f seconds", elapsed.Seconds())

}

func detrand() {
	// Will print out deterministic random shit
	fmt.Println("\nRANDOM SHIT GO!!!!!!!")

	// 1. Seed the rng
	rand.Seed(time.Now().UnixNano()) // This gets the current nanosecond to seed the rng

	// 2. RANDOM SHIT GO!!!!
	for i := 0; i < 7; i++ {
		fmt.Println("Random " + strconv.Itoa(rand.Intn(3478907)))
	}

}

func helloWorld() {
	// Gets user input then prints out greeting
	fmt.Println("\nPlease enter your name ...")

	// 1. Get input
	name := getText("Name: ")

	// 2. Greet
	fmt.Printf("Yo, nice to meet you %s! Great weather we're having, eh?", name)

}

func getText(ask string) string {
	// Gets keyboard input from a terminal user, and optionally prints out the ask before

	// 0. Print Ask
	fmt.Print(ask)

	// 1. Get data
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n') // Read up until a newline char

	if err != nil {
		fmt.Println("Unexpected error")
		return "" // End if unexpected error
	}

	// 2. Process text by removing carriage-return and newline chars
	text = strings.ReplaceAll(text, "\n", "")
	text = strings.ReplaceAll(text, "\r", "")

	// 3. Return clean text
	return text
}
