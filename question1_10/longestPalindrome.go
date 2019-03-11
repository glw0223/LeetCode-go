package question1_10

import "github.com/glw0223/LeetCode-go/util"

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
示例 1：
输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：
输入: "cbbd"
输出: "bb"
*/


func expandAroundCenter(s string, left,right int)(length int){
	for left>=0 && right<len(s) && s[left]==s[right] {
		left--
		right++
	}
	return right-left-1
}

func LongestPalindrome(s string) (result string) {
	if len(s)<1 {
		return ""
	}
	start,end := 0,0
	i:=0
	for ; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i + 1)
		len := util.Max(len1, len2)
		if len > end - start {
			start = i - (len - 1) / 2
			end = i + len / 2
		}
	}
	return s[start:end + 1]
}
