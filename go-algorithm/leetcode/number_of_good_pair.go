package leetcode

import "fmt"

/*
* https://leetcode.com/problems/number-of-good-pairs/
* Given an array of integers nums, return the number of good pairs.
* A pair (i, j) is called good if nums[i] == nums[j] and i < j.
 */
// example
/*
Input: nums = [1,2,3,1,1,3]
Output: 4
Explanation: There are 4 good pairs (0,3), (0,4), (3,4), (2,5) 0-indexed.
*/

func numIdenticicalPair2(nums []int) int {
	// O(N^2)
	// O(1)
	results := 0
	numberCount := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		count, isExist := numberCount[nums[i]]
		if isExist {
			results += count
			numberCount[nums[i]] = count + 1
		} else {
			numberCount[nums[i]] = 1
		}
	}

	return results
}
func numIdenticicalPair1(nums []int) int {
	// O(N^2)
	// O(1)
	results := 0

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				results += 1
			}
		}
	}

	return results
}

func TestNumIndenticalPair() {
	fmt.Println("Input: nums = [1,2,3,1,1,3]")
	input := []int{1, 2, 3, 1, 1, 3}
	fmt.Printf("Results 1: %v \n", numIdenticicalPair1(input))
	fmt.Printf("Results 2: %v \n", numIdenticicalPair2(input))

}
