package question1_10

import (
	"fmt"
	"github.com/glw0223/LeetCode-go/question1_10"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	str := "pwwkew"
	myLen := question1_10.LengthOfLongestSubstring(str)
	fmt.Println("str is ", str, " and longest substr len is ", myLen)
}
