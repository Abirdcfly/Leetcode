class Solution(object):
    def firstMissingPositive(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        x = 1
        while True:
            if x not in nums:
                return x
            x+=1
