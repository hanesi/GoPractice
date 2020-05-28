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
