package main

import (
	"ds/problems"
	"fmt"
)

func main() {
	// Problem 1: Find the repetetive characters in a string
	problems.FindRepetetiveValues("asdfasdfasd")
	/*Problem 2: Given a  list of numbers and a number k, return whether any two numbers from list would add up to k
	For Example given [10, 15, 3 , 7] and k  of 17 , return true since 10 + 7 is 17
	Bonus can you do this one pass.
	*/
	a := []int{10, 15, 3, 7}
	var k int
	k = 17
	ans := problems.FindKinArray(a, k)
	fmt.Println(ans)

}
