class Solution:
    def findMaxAverage(self, nums: List[int], k: int) -> float:
        maxSum = total = sum(nums[0:k])
        for i in range(k,len(nums)):
            total = total - nums[i-k] + nums[i]
            maxSum = max(maxSum, total)
        return float(maxSum) / float(k)