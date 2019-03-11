package question1_10

import (
	"fmt"
	"github.com/glw0223/LeetCode-go/question1_10"
	"testing"
)

func TestIsPalindrome(t *testing.T){
	result:=question1_10.IsPalindrome(121)
	fmt.Println(result)
	result=question1_10.IsPalindrome(-121)
	fmt.Println(result)
	result=question1_10.IsPalindrome(10)
	fmt.Println(result)
}