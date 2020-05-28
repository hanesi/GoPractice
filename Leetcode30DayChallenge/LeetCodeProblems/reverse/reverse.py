"""
This function takes an int and returns it reversed. If the int is bigger than
32 bits, it returns 0
"""
class Solution(object):
    def reverse(self, x):
        stringX = str(x)
        if x < 0:
            stringX = stringX[1:]
            if abs(int(stringX[::-1])) > ((1 << 31) - 1):
                return 0
            return 0 - int(stringX[::-1])
        if abs(int(stringX[::-1])) > 2147483647:
            return 0
        return int(stringX[::-1])

sol = Solution()
test = sol.reverse(123456789)
