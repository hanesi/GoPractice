"""
This function accepts a list of INTs and a target sum value.
It returns the indices of the two list values that combine
to result in the target sum. No repeated values and only
one possible target sum combination are possible.
"""
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        checkDict = {}
        for i,n in enumerate(nums):
            if n in checkDict:
                return [checkDict[n], i]
            checkDict[target - n] = i
        return []

numlist = [2, 7, 11, 15]
targ = 9

tmp = Solution()
ok = tmp.twoSum(numlist, targ)
