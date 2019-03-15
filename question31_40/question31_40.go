package question31_40

func A33_search(nums []int, target int) int {
	length:=len(nums)
	if length==0{
		return -1
	}
	left:=0;right:=length-1
	for left<=right {
		mid:=(left+right)/2
		if nums[mid] == target{
			return mid
		}
		if nums[left]<=nums[mid]{
			if nums[left]<=target && target<=nums[mid]{
				right=mid
			}else {
				left=mid+1
			}
		}else {
			if nums[mid]<=target && target<=nums[right] {
				left=mid
			}else {
				right=mid-1
			}
		}
	}
	return -1
}
