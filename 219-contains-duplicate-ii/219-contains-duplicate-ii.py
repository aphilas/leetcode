class Solution:
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:        
        hashMap = dict()
        index = 0

        while index < len(nums):
            current = nums[index]
            prevIndex = hashMap.get(current, None)

            if prevIndex is not None and (index - prevIndex) <= k:
                return True

            hashMap[current] = index
            index += 1

        return False
            