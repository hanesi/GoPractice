"""
Determine whether an integer is a palindrome.
An integer is a palindrome when it reads the same backward as forward.

Example 1:

Input: 121
Output: true
"""

class Solution:
    def isPalindrome(self, x: int) -> bool:
        if x < 0:
            return False
        xRev = str(x)[::-1]
        if int(xRev) == x:
            return True
        return False

tmp = Solution()
test = tmp.isPalindrome(123454321)
