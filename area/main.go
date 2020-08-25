package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}
type square struct {
	sideLength float64
}

func main() {
	// Create reader for reading from the terminal
	reader := bufio.NewReader(os.Stdin)

	// Get value of triangle base, trim '\n' and convert result to float64
	fmt.Print("Enter triangle base: ")
	tbi, _ := reader.ReadString('\n')
	tbp := strings.TrimSuffix(tbi, "\n")
	tb, _ := strconv.ParseFloat(tbp, 64)

	// Get value of triangle height, trim '\n' and convert result to float64
	fmt.Print("Enter triangle height: ")
	thi, _ := reader.ReadString('\n')
	thp := strings.TrimSuffix(thi, "\n")
	th, _ := strconv.ParseFloat(thp, 64)
	// Instantiate triangle struct with given base and height
	t := triangle{tb, th}

	// Get value of square side length, trim '\n' and convert result to float64
	fmt.Print("Enter square side length: ")
	ssli, _ := reader.ReadString('\n')
	sslp := strings.TrimSuffix(ssli, "\n")
	ssl, _ := strconv.ParseFloat(sslp, 64)
	// Instantiate square struct with given side length
	s := square{ssl}

	// Call interface functions to print area of square and triangle
	printArea(t)
	printArea(s)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
