class Solution(object):
    def firstMissingPositive(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        哈希思想
        """
        x = 1
        while True:
            if x not in nums:
                return x
            x+=1

class Solution(object):
    def firstMissingPositive(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        把[1,len(nums)]内的数按顺序排列，查找对应数字和下标的关系
        """
        new_nums = [-1]*(len(nums)+1)
        for i in nums:
            if i>0 and i<=len(nums):
                new_nums[i] = i
        new_nums = [i for i in new_nums if i != -1]
        for index, i in enumerate(new_nums):
            if index+1 != i:
                return index+1
        return len(new_nums)+1
