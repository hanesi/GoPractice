"""
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
"""


class Solution:
    def findMedianSortedArrays(self, nums1, nums2):
        a, b = sorted((nums1, nums2), key=len)
        m, n = len(a), len(b)
        after = (m + n - 1) // 2
        low, high = 0, m
        while low < high:
            i = (low + high) // 2
            if after-i-1 < 0 or a[i] >= b[after-i-1]:
                high = i
            else:
                low = i + 1
        i = low
        nextfew = sorted(a[i:i+2] + b[after-i:after-i+2])
        return (nextfew[0] + nextfew[1 - (m+n) % 2]) / 2.0


tmp = Solution()
print(tmp.findMedianSortedArrays([1, 3], [2]))
