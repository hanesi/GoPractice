"""
Given a string, determine if it is a palindrome,
considering only alphanumeric characters and ignoring cases.

Note: For the purpose of this problem, we
define empty string as valid palindrome.

Example 1:
Input: "A man, a plan, a canal: Panama"
Output: true
"""

import re


class Solution:
    def isPalindrome(self, s: str) -> bool:
        s = s.lower()
        s = re.sub(r'\W+', '', s)
        if s == s[::-1]:
            return True
        return False


tmp = Solution()
print(tmp.isPalindrome("A man, a plan, a canal: Panama"))
