package question31_40

/*
思路：1、先确定nums[mid]和旋转点的关系，也就是nums[left]或者nums[right]的关系
2、再判断单调的那个区间，使用常规二分法判断；else的逻辑就是移动另一边
3、最后收敛条件一定是在单调区间上
 */
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
				right=mid-1
			}else {
				left=mid+1
			}
		}else {
			if nums[mid]<=target && target<=nums[right] {
				left=mid+1
			}else {
				right=mid-1
			}
		}
	}
	return -1
}
