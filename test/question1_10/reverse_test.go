package question1_10

import (
	"fmt"
	"github.com/glw0223/LeetCode-go/question1_10"
	"testing"
)

func TestReverse(t *testing.T) {
	var result int
	result=question1_10.Myreverse(123)
	fmt.Println(result)
	result=question1_10.Myreverse(-123)
	fmt.Println(result)
	result=question1_10.Myreverse(120)
	fmt.Println(result)
}
