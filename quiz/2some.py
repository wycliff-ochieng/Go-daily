def twoSum(nums,target):
    nums_to_index = {}
    for i,num in enumerate(nums):
        complement = target - num
        if complement in nums_to_index:
            return [nums_to_index[complement],i]
        nums_to_index[num] = i
    return []

nums = [2,7,11,15]
target = 9       
processing = twoSum(nums,target)
print(processing)