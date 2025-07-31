class Solution:
    def maxArea(self, height: List[int]) -> int:
        left, right = 0, len(height)-1
        maxArea = 0
        while left < right:
            length = right-left
            currArea = min(height[left], height[right]) * length
            if currArea > maxArea:
                maxArea = currArea
            if height[left] < height[right]:
                left += 1
            else:
                right -= 1
        return maxArea