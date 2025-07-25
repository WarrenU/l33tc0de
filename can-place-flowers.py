class Solution:
    def canPlaceFlowers(self, fb: List[int], n: int) -> bool:
        ml = len(fb)
        if n == 0:
            return True
        for i in range(ml):
            lv = i == 0 or fb[i-1] == 0
            rv = i == ml - 1 or fb[i+1] == 0
            if lv and rv and fb[i] == 0:
                fb[i] = 1
                n -= 1
                if n == 0:
                    return True
                i += 1
        return n == 0