package test

import (
	"fmt"
	"github.com/glw0223/LeetCode-go/longestPalindrome"
	"testing"
)

func TestLongestPalindrome(t *testing.T){
	result:=longestPalindrome.LongestPalindrome("babad")
	fmt.Println(result)
}
