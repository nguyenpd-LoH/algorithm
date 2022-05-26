package leetcode

// https://leetcode.com/problems/water-bottles/
/**
There are numBottles water bottles that are initially full of water. You can exchange numExchange empty water bottles from the market with one full water bottle.
The operation of drinking a full water bottle turns it into an empty bottle.
Given the two integers numBottles and numExchange, return the maximum number of water bottles you can drink.
*/

// Input: numBottles = 9, numExchange = 3
// Output: 13
// Explanation: You can exchange 3 empty bottles to get 1 full water bottle.
// Number of water bottles you can drink: 9 + 3 + 1 = 13.

import "fmt"

func numWaterBottles(numBottles int, numExchange int) int {
	var results int
	var emptyBottles int

	// crink
	emptyBottles = numBottles
	results = numBottles

	for emptyBottles >= numExchange {
		numBottles = emptyBottles / numExchange
		emptyBottles = emptyBottles % numExchange

		emptyBottles += numBottles
		results += numBottles
	}
	return results
}

func TestWaterBottles() {
	fmt.Sprintf("Check with numBottles %v and numExchange  %v", 9, 3)
	results := numWaterBottles(9, 3)
	fmt.Println("Results:", results)
}
