package main

import "fmt"

func main() {
	// Declares a slice called myCourses
	myCourses := []string{"Docker", "Puppet", "Python"}
	// myCourses := make([]string, 5, 10)

	fmt.Printf("Length is: %d.\nCapacity is: %d", len(myCourses), cap(myCourses))
}
