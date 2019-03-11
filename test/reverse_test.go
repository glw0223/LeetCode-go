package test

import (
	"fmt"
	"github.com/glw0223/LeetCode-go/reverse"
	"testing"
)

func TestReverse(t *testing.T) {
	var result int
	result=reverse.Myreverse(123)
	fmt.Println(result)
	result=reverse.Myreverse(-123)
	fmt.Println(result)
	result=reverse.Myreverse(120)
	fmt.Println(result)
}
