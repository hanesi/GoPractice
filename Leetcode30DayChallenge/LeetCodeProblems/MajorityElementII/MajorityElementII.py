"""
Given an integer array of size n, find all elements
that appear more than ⌊ n/3 ⌋ times.

Note: The algorithm should run in linear time and in O(1) space.

Example 1:
Input: [3,2,3]
Output: [3]
"""


class Solution:
    def majorityElement(self, nums: list[int]) -> list[int]:
        ctDict = {}
        ctList = []
        for i in nums:
            ctDict[i] = ctDict.get(i, 0) + 1
        for k, v in ctDict.items():
            if v > len(nums)/3:
                ctList.append(k)
        return ctList


tmp = Solution()
print(tmp.majorityElement([3, 2, 3]))
