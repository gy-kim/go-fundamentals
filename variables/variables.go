package main

import "fmt"

// var (
// 	// name, course string
// 	// module       float64

// 	// name, course, module = "Nigel", "Docker Deep Dive", 3.2

// 	name   = "Nigel"
// 	course = "Docker Deep Dive"
// 	module = 3.2
// )

func main() {

	/*
		//// Pointers
		name := "Nigel"
		module := 3.2
		ptr := &module

		fmt.Println("Name is set to", name, "and is of type", reflect.TypeOf(name))
		fmt.Println("Module is set to", module, "and is of type", reflect.TypeOf(module))
		fmt.Println("Memory address of *module* variable is", ptr, "and the value of *module* is", *ptr)

	*/

	/*
		//// Passing by Value
		name := "Nigel"              // Name of subscriber
		course := "Docker Deep Dive" // Course being viewed

		fmt.Println("\nHi", name, "You're currently watching", course)

		changeCourseByValue(course)

		fmt.Println("\nYou are now watching course", course)
	*/

	/* */

	//// Passing by Reference
	name := "Nigel"              // Name of subscriber
	course := "Docker Deep Dive" // Course being viewed

	fmt.Println("\nHi", name, "You're currently watching", course)

	changeCourseByReference(&course)

	fmt.Println("\nYou are now watching course", course)

}

func changeCourseByReference(course *string) string {
	*course = "First Look: Native Docker Clustering"

	fmt.Println("\nTrying to change your course to", *course)

	return *course
}

func changeCourseByValue(course string) string {
	course = "First Look: Native Docker Clustering"

	fmt.Println("\nTrying to change your course to", course)

	return course
}
