class Solution:
    def missingNumber(self, nums: List[int]) -> int:
        mask = [False] * (len(nums)+1)

        for n in nums:
            mask[n] = True

        return mask.index(False)