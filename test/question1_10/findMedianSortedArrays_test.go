package question1_10

import (
	"fmt"
	"github.com/glw0223/LeetCode-go/question1_10"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	nums1:=[]int{1,2}
	nums2:=[]int{3,4}
	result:=question1_10.FindMedianSortedArrays(nums1,nums2)
	fmt.Println(result)
}
