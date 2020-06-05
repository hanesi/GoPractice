"""
Given an array of non-negative integers, you are
initially positioned at the first index of the array.

Each element in the array represents your maximum
jump length at that position.

Determine if you are able to reach the last index.

Example 1:
Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
"""


class Solution:
    def canJump(self, nums: list[int]) -> bool:
        lastPosition = len(nums) - 1
        for i in reversed(range(lastPosition)):
            if i + nums[i] >= lastPosition:
                lastPosition = i
        return lastPosition == 0


tmp = Solution()
print(tmp.canJump([2, 3, 1, 1, 4]))
