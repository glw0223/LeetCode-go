package test

import (
	"fmt"
	"github.com/glw0223/LeetCode-go/findMedianSortedArrays"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	nums1:=[]int{1,2}
	nums2:=[]int{3,4}
	result:=findMedianSortedArrays.FindMedianSortedArrays(nums1,nums2)
	fmt.Println(result)
}
