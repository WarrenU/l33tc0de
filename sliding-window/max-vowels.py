class Solution:
    def maxVowels(self, s: str, k: int) -> int:
        vowels = {'a', 'e', 'i', 'o', 'u'}
        count = sum(1 for c in s[:k] if c in vowels)
        maxV = count
        for i in range(k,len(s)):
            if s[i-k] in vowels:
                count -=1
            if s[i] in vowels:
                count += 1
            maxV = max(maxV, count)
        return maxV