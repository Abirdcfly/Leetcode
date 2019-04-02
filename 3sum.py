class Solution(object):
    def threeSum(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        ret = []
        nums.sort()
        for m in range(len(nums)-2):
            if m > 0 and nums[m] == nums[m-1]:
                continue
            l = m + 1
            r = len(nums) - 1
            while l < r:
                sums = nums[m] + nums[l] + nums[r]
                if sums == 0:
                    ret.append([nums[m],nums[r],nums[l]])
                    while nums[l] == nums[l+1] and l+1 < r:
                        l+=1
                    while nums[r] == nums[r-1]and l < r-1:
                        r-=1
                    l+=1
                    r-=1
                elif sums < 0:
                    l+=1
                elif sums > 0:
                    r-=1
        return ret
