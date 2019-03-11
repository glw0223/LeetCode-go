package question1_10

import (
	"fmt"
	"github.com/glw0223/LeetCode-go/question1_10"
	"testing"
)

func TestMyAtoi(t *testing.T){
	var result int
	result=question1_10.MyAtoi("42")
	fmt.Println(result)
	result=question1_10.MyAtoi("-42")
	fmt.Println(result)
	result=question1_10.MyAtoi("4193 with words")
	fmt.Println(result)
	result=question1_10.MyAtoi("-91283472332")
	fmt.Println(result)
}