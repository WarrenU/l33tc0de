class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        ip = 0
        for n in nums:
            if n != 0:
                nums[ip] = n
                ip += 1
        for i in range(ip, len(nums)):
            nums[i] = 0
        
        