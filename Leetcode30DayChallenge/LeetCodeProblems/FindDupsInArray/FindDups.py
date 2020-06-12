"""
Given an array of integers, 1 ≤ a[i] ≤ n (n = size of array),
some elements appear twice and others appear once.

Find all the elements that appear twice in this array.

Could you do it without extra space and in O(n) runtime?

Example:
Input:
[4,3,2,7,8,2,3,1]
Output:
[2,3]
"""


class Solution:
    def findDuplicates(self, nums: list[int]) -> list[int]:
        length = len(nums) + 1
        ctList = [0] * length
        reList = []
        for v in nums:
            ctList[v] += 1
        for i, x in enumerate(ctList):
            if x > 1:
                reList.append(i)
        return reList


tmp = Solution()
print(tmp.findDuplicates([4, 3, 2, 7, 8, 2, 3, 1]))
