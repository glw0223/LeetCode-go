package test

import (
	"fmt"
	"github.com/glw0223/LeetCode-go/myAtoi"
	"testing"
)

func TestMyAtoi(t *testing.T){
	var result int
	result=myAtoi.MyAtoi("42")
	fmt.Println(result)
	result=myAtoi.MyAtoi("-42")
	fmt.Println(result)
	result=myAtoi.MyAtoi("4193 with words")
	fmt.Println(result)
	result=myAtoi.MyAtoi("-91283472332")
	fmt.Println(result)
}