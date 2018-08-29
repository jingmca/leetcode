class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        for i,v in enumerate(nums):
            other = target - v
            if other in nums[i+1:]:
                return [i, nums.index(other, i+1)]

        return [-1,-1]