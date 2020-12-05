"""
Given two positive integers n and k.
A factor of an integer n is defined
as an integer i where n % i == 0.

Consider a list of all factors of n
sorted in ascending order, return the
kth factor in this list or return -1
if n has less than k factors.

Example 1:
Input: n = 12, k = 3
Output: 3
Explanation: Factors list is [1, 2, 3, 4, 6, 12], the 3rd factor is 3.
"""

class Solution:
    def kthFactor(self, n: int, k: int) -> int:
        factors = []
        for i in range(1, n+1):
            if n % i == 0:
                factors.append(i)
        print(factors)
        if len(factors) >= k:
            return factors[k-1]
        return -1

tmp = Solution()
tmp.kthFactor(12, 3)
