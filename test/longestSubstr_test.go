package test

import (
	"fmt"
	"github.com/glw0223/LeetCode-go/longestSubstr"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	str := "pwwkew"
	myLen := longestSubstr.LengthOfLongestSubstring(str)
	fmt.Println("str is ", str, " and longest substr len is ", myLen)
}
