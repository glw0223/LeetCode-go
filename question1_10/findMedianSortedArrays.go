package question1_10

import (
	"fmt"
	"github.com/glw0223/LeetCode-go/util"
)

/*
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5
 */
func FindMedianSortedArrays(nums1 []int, nums2 []int) (result float64) {
	m := len(nums1)
	n := len(nums2)
	if m > n { // to ensure m<=n
		temp := nums1
		nums1 = nums2
		nums2 = temp
		m = len(nums1)
		n = len(nums2)
	}
	fmt.Println("m:",m,"n:",n)
	fmt.Println("nums1:",nums1,"nums2:",nums2)
	iMin := 0; iMax := m; halfLen := (m + n + 1)>>1
	for iMin <= iMax {
		i := (iMin + iMax)>>1
		j := halfLen - i
		if i<iMax && nums2[j-1]>nums1[i] {
			iMin = i + 1 // i is too small
		}else if i>iMin && nums1[i-1]>nums2[j] {
			iMax = i - 1 // i is too big
		}else { // i is perfect
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			}else if j == 0 {
				maxLeft = nums1[i-1]
			}else {
				maxLeft = util.Max(nums1[i-1], nums2[j-1])
			}
			if  (m + n) % 2 == 1 {
				return float64(maxLeft)
			}

			minRight := 0;
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				minRight = util.Min(nums2[j], nums1[i])
			}

			return (float64(maxLeft + minRight))/2.0
		}
	}
	return 0.0
}
