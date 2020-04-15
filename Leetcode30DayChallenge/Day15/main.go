/*
Given an array nums of n integers where n > 1,
return an array output such that output[i] is
equal to the product of all the elements of nums except nums[i].

Example:

Input:  [1,2,3,4]
Output: [24,12,8,6]
Constraint: It's guaranteed that the product of the elements
of any prefix or suffix of the array (including the whole array)
fits in a 32 bit integer.

Note: Please solve it without division and in O(n).

Follow up:
Could you solve it with constant space complexity?
(The output array does not count as extra space for
the purpose of space complexity analysis.)
*/
package main

import "fmt"

func main() {
	testSlice := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(testSlice))
}

func productExceptSelf(nums []int) (answer []int) {
	answer = make([]int, len(nums))
	answer[0] = 1

	for i := 1; i < len(nums); i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}

	rightProduct := 1
	for i := len(nums) - 1; i > 0; i-- {
		answer[i] = answer[i] * rightProduct
		rightProduct *= nums[i]
	}
	answer[0] = rightProduct
	return
}
