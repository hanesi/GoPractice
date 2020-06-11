/*
There are two sorted arrays nums1 and
nums2 of size m and n respectively.

Find the median of the two sorted arrays.
The overall run time complexity should be
O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:
nums1 = [1, 3]
nums2 = [2]
The median is 2.0
*/
package main

import "fmt"

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}
	l := 0
	r := m
	for l <= r {
		i := (l + r) / 2
		j := (m+n+1)/2 - i
		if i > 0 && nums1[i-1] > nums2[j] {
			r = i - 1
		} else if i < m && nums2[j-1] > nums1[i] {
			l = i + 1
		} else {
			leftMax := 0
			if i == 0 {
				leftMax = nums2[j-1]
			} else if j == 0 {
				leftMax = nums1[i-1]
			} else {
				leftMax = max(nums2[j-1], nums1[i-1])
			}
			if (m+n)&1 == 1 {
				return float64(leftMax)
			}
			rightMin := 0
			if i == m {
				rightMin = nums2[j]
			} else if j == n {
				rightMin = nums1[i]
			} else {
				rightMin = min(nums2[j], nums1[i])
			}
			return float64(leftMax+rightMin) / 2
		}
	}
	return 0.0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
