package main

import "fmt"

func main() {
	/*
		//// Slice syntax
		myCourses := []string{"Docker", "Puppet", "Python"}
		// myCourses := make([]string, 5, 10)

		fmt.Printf("Length is: %d.\nCapacity is: %d", len(myCourses), cap(myCourses))
	*/

	/*
		//// Getting Under the Hood
		mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		fmt.Println(mySlice[4])

		mySlice[1] = 0
		fmt.Println(mySlice)

		sliceOfSlice := mySlice[2:5]
		fmt.Println(sliceOfSlice)

	*/

	/*
		//// Append
		mySlice := make([]int, 1, 4)
		fmt.Printf("Length is: %d Capacity is: %d", len(mySlice), cap(mySlice))

		for i := 1; i < 17; i++ {
			mySlice = append(mySlice, i)
			fmt.Printf("\nCapacity is: %d", cap(mySlice))
		}
	*/

	/* */
	//// Miscellaneous
	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice)

	for _, i := range mySlice {
		fmt.Println("for range loop:", i)
	}

	newSlice := []int{10, 20, 30}
	mySlice = append(mySlice, newSlice...)
	fmt.Println(mySlice)

}
