package question11_20

import (
	"github.com/glw0223/LeetCode-go/util"
)

/*
给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。
示例：
输入: [1,8,6,2,5,4,8,3,7]
输出: 49

 */

func MaxArea(height []int) (result int) {
	length:=len(height)
	for i:=0;i<length;i++{
		for j:=i+1;j<length;j++{
			result = util.Max(result, util.Min(height[i],height[j])*(j-i))
		}
	}
	return
}
func MaxArea1(height []int) (result int) {
	length:=len(height)
	left:=0
	right:=length-1
	for left<right{
		result = util.Max(result, util.Min(height[left],height[right])*(right-left))
		if height[left]<height[right]{
			left++
		}else {
			right--
		}
	}
	return
}


