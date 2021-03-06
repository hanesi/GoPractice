"""
Given an array nums and a value val, remove
all instances of that value in-place and
return the new length.

Do not allocate extra space for another
array, you must do this by modifying the
input array in-place with O(1) extra memory.

The order of elements can be changed. It
doesn't matter what you leave beyond the
new length.

Example 1:
Given nums = [3,2,2,3], val = 3,

Your function should return length = 2,
with the first two elements of nums being 2.

It doesn't matter what you leave beyond the
returned length.
"""


class Solution:
    def removeElement(self, nums: list[int], val: int) -> int:
        length = len(nums) - 1
        while length >= 0:
            if nums[length] == val:
                nums.remove(nums[length])
            length -= 1
        return len(nums)

tmp = Solution()
print(tmp.removeElement([3,2,2,3]))
